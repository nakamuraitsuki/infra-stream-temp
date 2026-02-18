import { useState } from "react";
import { useAuth } from "../../context/AuthContext"
import { useServices } from "../../context/ServiceContext";
import type { VideoId, VideoTag } from "../../domain/video/video.model";

export type CreateVideoResult =
  | { type: "complete", videoId: VideoId }
  | { type: "upload_failed", videoId: VideoId, error: string }
  | { type: "create_failed", error: string }
  | { type: "unauthenticated", error: string };

export const useCreateVideo = () => {
  const { session } = useAuth();
  const { videoRepo } = useServices();

  const [createLoading, setCreateLoading] = useState(false);
  const [uploadLoading, setUploadLoading] = useState(false);
  const loading = createLoading || uploadLoading;

  const createVideo = async (
    title: string,
    description: string,
    tags: VideoTag[],
    file: File
  ): Promise<CreateVideoResult> => {
    if (!session) {
      return { type: "unauthenticated", error: "User is not authenticated" };
    }

    setCreateLoading(true);
    let videoId: VideoId;
    try {
      const createRes = await videoRepo.create(title, description, tags);
      if (!createRes.success) {
        throw new Error(createRes.error);
      }
      videoId = createRes.data.id;
    } catch (e: any) {
      return { type: "create_failed", error: e.message };
    } finally {
      setCreateLoading(false);
    }

    setUploadLoading(true);
    try {
      const uploadRes = await videoRepo.uploadSource(videoId, file);
      if (!uploadRes.success) {
        throw new Error(uploadRes.error);
      }
      return { type: "complete", videoId };
    } catch (e: any) {
      return { type: "upload_failed", videoId, error: e.message };
    } finally {
      setUploadLoading(false);
    }
  };

  const retryUpload = async (videoId: VideoId, file: File): Promise<CreateVideoResult> => {
    try {
      setUploadLoading(true);

      const uploadRes = await videoRepo.uploadSource(videoId, file);
      if (!uploadRes.success) {
        throw new Error(uploadRes.error);
      }
      return { type: "complete", videoId };
    } catch (e: any) {
      return { type: "upload_failed", videoId, error: e.message };
    } finally {
      setUploadLoading(false);
    }
  }

  return { 
    createVideo, 
    retryUpload, 
    loading, 
    createLoading, 
    uploadLoading 
  };
}