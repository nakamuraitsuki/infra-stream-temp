import { useNavigate } from "react-router";
import { useAuth } from "../../context/AuthContext";
import { useMyVideos } from "../../hooks/video/use-my-videos"
import { useEffect, useState } from "react";
import { type Video } from "../../domain/video/video.model";

const MY_PAGE_LIMIT = 10;

export const MyPage = () => {
  const navigate = useNavigate();
  const { session } = useAuth();

  const { fetch, loading: listLoading } = useMyVideos();
  const [videos, setVideos] = useState<Video[]>([]);

  const loadList = async () => {
    const res = await fetch(MY_PAGE_LIMIT);
    if (res.type === "success") {
      setVideos(res.videos);
    }
  };

  useEffect(() => {
    if (session.status === "authenticated") {
      loadList();
    } else if (session.status === "unauthenticated") {
      navigate("/");
    }
    // NOTE: session が loading のときはスルーする
  }, [session.status]);

  if (session.status === "unauthenticated") return null;

  return (
    <div style={{ padding: "20px", maxWidth: "800px", margin: "0 auto" }}>
      <h2>My Page</h2>
      <p>Welcome, {session.user?.name}!</p>

      <hr />
      <p>アップロード</p>
      <button onClick={() => navigate("/upload")}>Upload New Video</button>
      <section>
        <h3>My Videos</h3>
        {listLoading ? (
          <p>Loading videos...</p>
        ) : videos.length === 0 ? (
          <p> No videos found.</p>
        ) : (
          <div style={{ display: "grid", gap: "10px" }}>
            {videos.map((video) => (
              <div
                key={video.id}
                onClick={() => navigate(`/video/${video.id}`)}
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
                <span>{video.tags.join(", ")}</span>
              </div>
            ))}
          </div>
        )}
      </section>
    </div>
  )
}