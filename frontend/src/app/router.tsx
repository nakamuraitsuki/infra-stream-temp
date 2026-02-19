import { createBrowserRouter, Navigate } from "react-router"
import { HomePage } from "../pages/home-page/HomePage"
import { MyPage } from "../pages/my-page/MyPage"
import { VideoUploadPage } from "../pages/video-upload-page/VideoUploadPage"
import { VideoPlayPage } from "../pages/video-play-page/VideoPlayPage"

export const router = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/my-page",
    element: <MyPage />,
  },
  {
    path: "/upload",
    element: <VideoUploadPage />,
  },
  {
    path: "/video/:videoId",
    element: <VideoPlayPage />,
  },
  {
    path: "*",
    element: <Navigate to="/" replace />,
  },
])
