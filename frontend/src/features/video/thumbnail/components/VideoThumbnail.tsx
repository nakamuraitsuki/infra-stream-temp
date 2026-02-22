import type { VideoId } from "@/domain/video/video.model"
import styles from "./VideoThumbnail.module.css";

type Props = {
  videoId: VideoId;
  title: string;
  status?: string;
}

export const VideoThumbnail = ({ title, status = "ready" }: Props) => {
  const overlayClass = status === "ready" ? styles.overlayNormal : styles.overlayUnavailable;
  return (
    <div className={styles.thumbnail}>
      <img src={`/thumbnail/cover1.jpg`} alt={title} className={styles.image} />
      <div className={overlayClass} />
      <div className={styles.title}>{title}</div>
    </div>
  );
}