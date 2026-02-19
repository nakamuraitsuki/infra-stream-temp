import { useNavigate } from "react-router";
import { useLogin } from "../../hooks/auth/use-login"
import { useState } from "react";

export const HomePage = () => {
  const { login } = useLogin();
  const navigator = useNavigate();

  const [error, setError] = useState<string | null>(null);

  const onLoginClick = async () => {
    setError(null);
    // ダミーユーザーでログイン
    const res = await login("dummy_user", "dummy_password");
    if (!res.success) {
      setError("ログインに失敗しました");
      return;
    }
    navigator("/my-page");
  }

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h2>Home</h2>
      <p>デモ用にダミーログインを行う</p>
      <button onClick={onLoginClick} style={{ padding: "10px 20px" }}>
        Login( Set Dummy User )
      </button>
      {error && <p style={{ color: "red" }}>{error}</p>}
    </div>
  )
}
