import type { Video, VideoId } from "@/domain/video/video.model"
import styles from "./PublicVideoCard.module.css";
import { VideoThumbnail } from "../../thumbnail";

type Props = {
  video: Video;
  onClick?: (videoId: VideoId) => void;
};

export const PublicVideoCard = ({ video, onClick }: Props) => {
  return (
    <div
      className={styles.card}
      onClick={() => onClick?.(video.id)}
    >
      <VideoThumbnail
        videoId={video.id}
        title={video.title}
      />
      <div className={styles.cardBody}>
        <h4 className={styles.title}>{video.title}</h4>
        <p className={styles.description}>{video.description}</p>
      </div>
    </div>
  );
};