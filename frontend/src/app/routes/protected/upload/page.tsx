import { useNavigate } from "react-router";
import { VideoUploadForm } from "@/features/video/create/VideoUploadForm";

export const VideoUploadPage = () => {
  const navigate = useNavigate();

  return (
    <VideoUploadForm
      onBack={() => navigate(-1)}
      onSuccess={() => navigate("/my")}
    />
  );
};
