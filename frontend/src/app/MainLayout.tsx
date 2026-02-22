import { Link, Outlet } from "react-router"
import styles from "./MainLayout.module.css"
import { useAuth } from "../context/AuthContext"

export const MainLayout = () => {
  const { session } = useAuth();
  return (
    <div className={styles.wrapper}>
      <header className={styles.header}>
        <div className={styles.inner}>
          <Link to="/" className={styles.logo}>
            Infra Stream
          </Link>
          <nav className={styles.nav}>
            <Link to="/" className={styles.navLink}>Home</Link>
            { session.status === "authenticated" ? (
              <Link to="/upload" className={styles.navLink}>Library</Link>
            ) : session.status === "unauthenticated" ? (
              <Link to="/login">Login</Link>
            ) : (
              <p>Loading...</p>
            )}
          </nav>
        </div>
      </header>
      <main>
        <Outlet /> {/* children routes render here */}
      </main>
    </div>
  )
}