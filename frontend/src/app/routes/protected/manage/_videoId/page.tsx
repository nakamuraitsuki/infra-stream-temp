import { useNavigate, useLocation, useParams } from "react-router";
import { useAuth } from "@/context/AuthContext";
import { useEffect } from "react";
import { VideoEditPanel } from "@/features/video/edit";
import type { Video } from "@/domain/video/video.model";

export const VideoManagePage = () => {
  const navigate = useNavigate();
  const { session } = useAuth();
  const { videoId } = useParams<{ videoId: string }>();
  const location = useLocation();
  const video = location.state?.video as Video | undefined;

  useEffect(() => {
    if (session.status === "unauthenticated") {
      navigate("/");
    }
  }, [session.status, navigate]);

  if (session.status !== "authenticated") return null;

  if (!video || video.id !== videoId) {
    return (
      <div style={{ padding: "20px" }}>
        <p>動画情報が見つかりませんでした。</p>
        <button onClick={() => navigate("/my-page")}>マイページに戻る</button>
      </div>
    );
  }

  return (
    <div style={{ padding: "20px", maxWidth: "800px", margin: "0 auto" }}>
      <button onClick={() => navigate("/my-page")} style={{ marginBottom: "16px" }}>
        ← マイページに戻る
      </button>
      <h2 style={{ marginBottom: "24px" }}>{video.title}</h2>
      <VideoEditPanel video={video} />
    </div>
  );
};
