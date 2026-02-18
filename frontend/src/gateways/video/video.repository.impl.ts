import axios from "axios";
import { apiClient } from "../../api/client";
import type { VideoId, Video, Tag } from "../../domain/video/video.model";
import type { GetPlaybackInfoResponse, IVideoRepository, VideoError } from "../../domain/video/video.repository";
import { failure, success, type Result } from "../../domain/core/result";

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
      const { data } = await apiClient.get(`api/videos/${id}`);
      return success(data);
    } catch (error) {
      return failure(this.handleError(error));
    }
  }

  async findPublicVideos(): Promise<Video[]> {
    const { data } = await apiClient.get<Video[]>("api/videos");
    return data;
  }

  async findByTag(tag: Tag): Promise<Video[]> {
    const { data } = await apiClient.get<Video[]>("/api/videos/search", {
      params: { tag },
    })
    return data;
  }

  async findMyVideos(): Promise<Video[]> {
    const { data } = await apiClient.get<Video[]>("/videos/mine");
    return data;
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

  async create(title: string, description: string, tags: Tag[]): Promise<Result<Video, VideoError>> {
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