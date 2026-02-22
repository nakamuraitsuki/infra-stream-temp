import { useMemo, useState } from "react";
import { useAuth } from "../../context/AuthContext";
import { useServices } from "../../context/ServiceContext";
import { getPublicVideos } from "../../application/video/getPublicVideos.usecase";

export const usePublicVideos = () => {
  const { session } = useAuth();
  const { videoRepo } = useServices();

  const [loading, setLoading] = useState<boolean>(false);

  const execute = useMemo(
    () => getPublicVideos({ videoRepo, session }),
    [videoRepo, session]
  );

  const fetch = async (limit: number) => {
    setLoading(true);
    const res = await execute.execute(limit);
    setLoading(false);
    return res;
  }

  return { fetch, loading };
};