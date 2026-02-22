import { useNavigate } from "react-router";
import { useAuth } from "@/context/AuthContext";
import { useEffect } from "react";
import { MyVideosSection } from "@/features/video/list-mine/MyVideoSection";

const MY_PAGE_LIMIT = 10;

export const MyPage = () => {
  const navigate = useNavigate();
  const { session } = useAuth();

  useEffect(() => {
    if (session.status === "unauthenticated") {
      navigate("/");
    }
  }, [session.status, navigate]);

  if (session.status !== "authenticated") return null;

  return (
    <div style={{ padding: "20px", maxWidth: "800px", margin: "0 auto" }}>
      <h2>My Page</h2>
      <p>Welcome, {session.user?.name}!</p>

      <hr />

      <p>アップロード</p>
      <button onClick={() => navigate("/upload")}>
        Upload New Video
      </button>

      <MyVideosSection
        limit={MY_PAGE_LIMIT}
        onSelect={(id) => navigate(`/video/${id}`)}
      />
    </div>
  );
};
