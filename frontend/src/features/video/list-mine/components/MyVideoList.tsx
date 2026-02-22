import type { Video, VideoId } from "@/domain/video/video.model";
import { useMyVideosQuery } from "../hooks/useMyVideosQuery";
import styles from "./MyVideoList.module.css";
import { MyVideoCard } from "./MyVideoCard";

type Props = {
  limit?: number;
  onSelect?: (videoId: VideoId) => void;
}

export const MyVideoList = ({ limit = 20, onSelect }: Props) => {
  const { data, isLoading, error } = useMyVideosQuery(limit);

  if (isLoading) return <p>loading my videos...</p>
  if (error) return <p>{error.message}</p>
  if (!data || data.videos.length === 0) return <p>No videos available.</p>

  return (
    <div className={styles.grid}>
      {data.videos.map((video: Video) => (
        <MyVideoCard
          key={video.id}
          video={video}
          onClick={onSelect}
        />
      ))}
    </div>
  );
};