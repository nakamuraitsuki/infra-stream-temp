import { useNavigate, useParams } from "react-router";
import { useVideoPlayer } from "@/features/video/play/use-video-player";
import { useEffect, useRef, useState } from "react";
import type { PlaybackDetail } from "../../../application/video/getPlaybackDetail.usecase";
import Hls from "hls.js";

export const VideoPlayPage = () => {
  const { videoId } = useParams<{ videoId: string }>();
  const navigate = useNavigate();
  const videoRef = useRef<HTMLVideoElement>(null);

  const { load, loading } = useVideoPlayer();

  const [detail, setDetail] = useState<PlaybackDetail | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (videoId) {
      load(videoId).then(res => {
        if (res.type === "success") setDetail(res.detail);
        else setError(res.error);
      });
    }
  }, [videoId, load]);

  useEffect(() => {
    const video = videoRef.current;
    if (!video || !detail) return;

    if (!Hls.isSupported()) {
      setError("HLS is not supported in this browser");
      return;
    }

    const { url: playbackUrl } = detail;
    const hls = new Hls();
    hls.loadSource(playbackUrl);
    hls.attachMedia(video);

    return () => {
      hls.destroy();
    };
  }, [detail]);

  if (loading) return <div style={{ padding: "20px" }}>Loading video...</div>;
  if (error) return (
    <div style={{ padding: "20px", color: "red" }}>
      <p>Error: {error}</p>
      <button onClick={() => navigate("/my-page")}>Back to My Page</button>
    </div>
  );

  return (
    <div style={{ padding: "20px" }}>
      <h2>{detail?.video?.title}</h2>
      <div style={{ maxWidth: "800px", background: "#000" }}>
        <video
          ref={videoRef}
          controls
          style={{ width: "100%", height: "100%" }}
        />
      </div>
      <div style={{ marginTop: "10px" }}>
        <p>{detail?.video?.description}</p>
      </div>
    </div>
  )
}