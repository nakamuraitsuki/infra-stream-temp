import { useEffect, useState } from "react";
import { type Video } from "@/domain/video/video.model";
import { useMyVideos } from "./use-my-videos";

type Props = {
  limit: number;
  onSelect: (videoId: string) => void;
};

export const MyVideosSection = ({ limit, onSelect }: Props) => {
  const { fetch, loading } = useMyVideos();

  const [videos, setVideos] = useState<Video[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const load = async () => {
      const res = await fetch(limit);

      if (res.type === "success") {
        setVideos(res.videos);
      } else {
        setError("Failed to load videos");
      }
    };

    load();
  }, [limit]);

  return (
    <section>
      <h3>My Videos</h3>

      {loading ? (
        <p>Loading videos...</p>
      ) : error ? (
        <p>{error}</p>
      ) : videos.length === 0 ? (
        <p>No videos found.</p>
      ) : (
        <div style={{ display: "grid", gap: "10px" }}>
          {videos.map((video) => (
            <div
              key={video.id}
              onClick={() => onSelect(video.id)}
              style={{
                padding: "15px",
                border: "1px solid #ccc",
                borderRadius: "8px",
                cursor: "pointer",
                display: "flex",
                justifyContent: "space-between"
              }}
            >
              <strong>{video.title}</strong>
              <span>{video.status}</span>
              <span>{video.visibility}</span>
              <span>{video.tags.join(", ")}</span>
            </div>
          ))}
        </div>
      )}
    </section>
  );
};
