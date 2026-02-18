import axios from "axios";
import { apiClient } from "../../api/client";
import type { VideoId, Video, VideoTag, VideoStatus, VideoVisibility } from "../../domain/video/video.model";
import type { GetPlaybackInfoResponse, IVideoRepository, VideoError } from "../../domain/video/video.repository";
import { failure, success, type Result } from "../../domain/core/result";
import { parseFindByTagResponse, parseFindMyVideosResponse, parseListPublicResponse } from "./video.dto";

export class VideoRepositoryImpl implements IVideoRepository {
  // helper method to map axios errors to VideoError
  private handleError(error: unknown): VideoError {
    if (axios.isAxiosError(error)) {
      const status = error.response?.status;
      if (status === 404) return "NOT_FOUND";
      if (status === 401) return "UNAUTHORIZED";
      if (status === 400) return "VALIDATION_ERROR";
      if (status && status >= 500) return "SERVER_ERROR";
    }
    return "UNKNOWN_ERROR";
  }

  async findByID(id: VideoId): Promise<Result<Video, VideoError>> {
    try {
      const { data } = await apiClient.get(`/api/videos/${id}`);
      return success(data);
    } catch (error) {
      return failure(this.handleError(error));
    }
  }

  async findPublicVideos(): Promise<Video[]> {
    try {
      const { data } = await apiClient.get("/api/videos");

      const dto = parseListPublicResponse(data);
      return dto.items.map((item): Video => ({
        id: item.id,
        ownerId: item.ownerId,
        title: item.title,
        description: item.description,
        tags: item.tags as VideoTag[],
        createdAt: new Date(item.createdAt),
      }))
    } catch (error) {
      console.error("Failed to fetch public videos:", error);
      return [];
    }
  }

  async findByTag(tag: VideoTag): Promise<Video[]> {
    try {
      const { data } = await apiClient.get("/api/videos/search", {
        params: { tag },
      })
      const dto = parseFindByTagResponse(data);
      return dto.items.map((item): Video => ({
        id: item.id,
        ownerId: item.ownerId,
        title: item.title,
        description: item.description,
        tags: item.tags as VideoTag[],
        createdAt: new Date(item.createdAt),
      }));
    } catch (error) {
      console.error(`Failed to search videos by tag "${tag}":`, error);
      return [];
    }
  }

  async findMyVideos(): Promise<Video[]> {
    try {
      const { data } = await apiClient.get("/api/videos/mine");
      const dto = parseFindMyVideosResponse(data);
      return dto.items.map((item): Video => ({
        id: item.id,
        title: item.title,
        description: item.description,
        tags: item.tags as VideoTag[],
        status: item.status as VideoStatus,
        visibility: item.visibility as VideoVisibility,
        createdAt: new Date(item.createdAt),
    }));
  } catch (error) {
      console.error("Failed to fetch my videos:", error);
      return [];
    }
  }

  async getPlaybackInfo(id: VideoId): Promise<Result<GetPlaybackInfoResponse, VideoError>> {
    try {
      const { data } = await apiClient.get<GetPlaybackInfoResponse>(
        `/api/videos/${id}/playback-info`,
      );
      return success(data);
    } catch (error) {
      return failure(this.handleError(error));
    }
  }

  async create(title: string, description: string, tags: VideoTag[]): Promise<Result<Video, VideoError>> {
    try {
      const { data } = await apiClient.post<Video>("/api/videos", {
        title,
        description,
        tags,
      });
      return success(data);
    } catch (error) {
      return failure(this.handleError(error));
    }
  }

  async uploadSource(id: VideoId, file: File): Promise<Result<void, VideoError>> {
    const formdata = new FormData();
    formdata.append("file", file);

    try {
      await apiClient.post(`/api/videos/${id}/source`, formdata, {
        timeout: 0, // アップロードは時間がかかる可能性があるため、タイムアウトを無効化
      });
      return success(undefined);
    } catch (error) {
      return failure(this.handleError(error));
    }
  }
}