import { useMemo, useState } from "react";
import { useAuth } from "../../context/AuthContext"
import { useServices } from "../../context/ServiceContext";
import type { VideoTag } from "../../domain/video/video.model";
import { createVideoMetaUseCase } from "../../application/video/createVideoMeta.usecase";
import { uploadVideoSource } from "../../application/video/uploadVideoSource.usecase";

export type CreateVideoResult =
  | { type: "success"; videoId: string }
  | { type: "create_failed"; error: string }
  | { type: "upload_failed"; error: string }
  | { type: "unauthenticated"; error: string };

export const useCreateVideo = () => {
  const { session } = useAuth();
  const { videoRepo } = useServices();

  const [createLoading, setCreateLoading] = useState<boolean>(false);
  const [uploadLoading, setUploadLoading] = useState<boolean>(false);

  const createMeta = useMemo(
    () => createVideoMetaUseCase({ session, videoRepo }),
    [videoRepo, session]
  );

  const uploadSource = useMemo(
    () => uploadVideoSource({ videoRepo }),
    [videoRepo]
  );

  const createVideo = async (
    title: string,
    description: string,
    tags: VideoTag[],
    file: File
  ): Promise<CreateVideoResult> => {
    setCreateLoading(true);
    const createRes = await createMeta.execute({ title, description, tags });
    setCreateLoading(false);

    if (createRes.type === "unauthenticated" || createRes.type === "create_failed") {
      return createRes;
    }

    const videoId = createRes.videoId;

    setUploadLoading(true);
    const uploadRes = await uploadSource.execute({ videoId, file });
    setUploadLoading(false);

    if (uploadRes.type === "upload_failed") {
      return uploadRes;
    }

    return { type: "success", videoId };
  };

  return {
    createVideo,
    createLoading,
    uploadLoading,
    loading: createLoading || uploadLoading
  };
};