import type { Tag, Video, VideoId } from "./video.model";

export interface GetPlaybackInfoResponse {
  playbackUrl: string; // Master playlist URL (e.g., HLS)
  mimeType: string;
}

export interface IVideoRepository {
  findByID(id: VideoId): Promise<Video | null>;
  findPublicVideos(): Promise<Video[]>;
  findByTag(tag: Tag): Promise<Video[]>;
  getPlaybackInfo(id: VideoId): Promise<GetPlaybackInfoResponse>;
  create(
    title: string,
    description: string,
    tags: Tag[],
  ): Promise<Video>;
  uploadSource(id: VideoId, file: File): Promise<void>;
  findMyVideos(): Promise<Video[]>;
}