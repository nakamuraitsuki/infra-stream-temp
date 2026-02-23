import { useAuth } from "@/context/AuthContext";
import { useEffect } from "react"
import { useNavigate } from "react-router";

export const LoginPage = () => {
  const { login } = useAuth();
  const navigate = useNavigate();
  // 現在はDummyなので、ログイン処理をしてすぐにリダイレクトする
  useEffect(() => {
    console.log("Logging in...");
    login("", "");
    navigate("/");
  }, [])
  return null;
}