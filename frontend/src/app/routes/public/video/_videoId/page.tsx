import { useNavigate, useParams } from "react-router";
import { VideoPlayer } from "@/features/video/play/VideoPlayer";

export const VideoPlayPage = () => {
  const { videoId } = useParams<{ videoId: string }>();
  const navigate = useNavigate();

  if (!videoId) return null;

  return (
    <VideoPlayer
      videoId={videoId}
      onErrorBack={() => navigate("/my")}
    />
  );
};
