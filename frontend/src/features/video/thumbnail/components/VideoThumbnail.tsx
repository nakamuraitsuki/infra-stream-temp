import type { VideoId } from "@/domain/video/video.model"
import styles from "./VideoThumbnail.module.css";
import { hashStringToIndex } from "../utils/hash";

type Props = {
  videoId: VideoId;
  title: string;
}

const gradientClasses = [
  styles.grandient0,
  styles.grandient1,
  styles.grandient2,
  styles.grandient3,
]

export const VideoThumbnail = ({ videoId, title }: Props) => {
  const index = hashStringToIndex(videoId as string, gradientClasses.length);
  const backgroundClass = gradientClasses[index];

  return (
    <div className={`${styles.thumbnail} ${backgroundClass}`}>
      <div className={styles.overlay} />
      <div className={styles.title}>{title}</div>
    </div>
  );
}