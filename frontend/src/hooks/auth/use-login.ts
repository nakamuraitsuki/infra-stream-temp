import { useAuth } from "../../context/AuthContext"

export const useLogin = () => {
  const { login } = useAuth();

  return { login }
}