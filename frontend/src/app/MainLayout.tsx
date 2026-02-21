import { Link, Outlet } from "react-router"
import styles from "./MainLayout.module.css"

export const MainLayout = () => {
  return (
    <div className={styles.wrapper}>
      <header className={styles.header}>
        <div className={styles.inner}>
          <Link to="/" className={styles.logo}>
            Infra Stream
          </Link>
          <nav className={styles.nav}>
            <Link to="/" className={styles.navLink}>Home</Link>
            <Link to="/my-page" className={styles.navLink}>Library</Link>
          </nav>
        </div>
      </header>
      <main>
        <Outlet /> {/* children routes render here */}
      </main>
    </div>
  )
}