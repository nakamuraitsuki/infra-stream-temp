import type { VideoId } from "@/domain/video/video.model"
import styles from "./VideoThumbnail.module.css";
import { hashStringToIndex } from "../utils/hash";

type Props = {
  videoId: VideoId;
  title: string;
  status?: string;
}

const gradientClasses = [
  styles.gradient0,
  styles.gradient1,
  styles.gradient2,
  styles.gradient3,
]

export const VideoThumbnail = ({ videoId, title, status = "ready" }: Props) => {
  const index = hashStringToIndex(videoId as string, gradientClasses.length);
  const backgroundClass = gradientClasses[index];

  const overlayClass = status === "ready" ? styles.overlayNormal : styles.overlayUnavailable;
  return (
    <div className={`${styles.thumbnail} ${backgroundClass}`}>
      <div className={overlayClass} />
      <div className={styles.title}>{title}</div>
    </div>
  );
}