import { useMemo, useState } from "react";
import { useAuth } from "../../context/AuthContext";
import { useServices } from "../../context/ServiceContext";
import { getMyVideos } from "../../application/video/getMyVideos";

export const useMyVideos = () => {
  const { session } = useAuth();
  const { videoRepo } = useServices();

  const [loading, setLoading] = useState(false);

  const execute = useMemo(
    () => getMyVideos({ videoRepo, session }),
    [videoRepo, session]
  );

  const fetch = async () => {
    setLoading(true);
    const res = await execute.execute();
    setLoading(false);
    return res;
  };

  return { fetch, loading };
};