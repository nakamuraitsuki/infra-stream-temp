import { useIsFetching } from "@tanstack/react-query"
import styles from "./SpinnerOverlay.module.css";

export const SpinnerOverlay = () => {
  const isFetching = useIsFetching();

  if (!isFetching) {
    return null;
  }

  return (
    <div className={styles.overlay}>
      <div className={styles.spinner}></div>
    </div>
  )
}