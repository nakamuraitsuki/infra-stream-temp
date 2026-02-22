import type { Video, VideoId } from "@/domain/video/video.model"
import { VideoThumbnail } from "../../thumbnail";
import styles from "./MyVideoCard.module.css";

type Props = {
  video: Video;
  onClick?: (videoId: VideoId) => void;
};

export const MyVideoCard = ({ video, onClick }: Props) => {
  return (
    <div
      className={styles.card}
      onClick={() => onClick?.(video.id)}
    >
      <div className={styles.thumbnail}>
        <VideoThumbnail
          videoId={video.id}
          title={video.title}
        />
      </div>
      <div className={styles.cardBody}>
        {/*内容*/}
      </div>
    </div>
  );
};