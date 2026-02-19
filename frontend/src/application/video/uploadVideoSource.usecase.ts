import type { VideoId } from "../../domain/video/video.model";
import type { IVideoRepository } from "../../domain/video/video.repository";

export type UploadVideoSourceResult =
  | { type: "success" }
  | { type: "upload_failed"; error: string };

export type UploadVideoSourceDeps = {
  videoRepo: IVideoRepository;
}

export type UploadVideoSourceParams = {
  videoId: VideoId;
  file: File;
}

export interface IUploadVideoSourceUseCase {
  execute(params: UploadVideoSourceParams): Promise<UploadVideoSourceResult>;
}

export const uploadVideoSource =
  ({ videoRepo }: UploadVideoSourceDeps): IUploadVideoSourceUseCase => ({
    execute: async ({ videoId, file }) => {
      const res = await videoRepo.uploadSource(videoId, file);

      if (!res.success) {
        return { type: "upload_failed", error: res.error };
      }

      return { type: "success" };
    }
  });
