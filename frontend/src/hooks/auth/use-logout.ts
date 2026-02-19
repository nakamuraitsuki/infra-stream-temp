import { useAuth } from "../../context/AuthContext";

export const useLogout = () => {
  const { logout } = useAuth();

  return { logout }
}