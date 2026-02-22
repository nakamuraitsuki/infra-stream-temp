import { createBrowserRouter, Navigate } from "react-router"
import { VideoUploadPage } from "../features/video/create/VideoUploadPage"
import { VideoPlayPage } from "../features/video/play/VideoPlayPage"
import { MainLayout } from "./MainLayout"
import { HomePage, MyPage } from "./routes"

export const router = createBrowserRouter([
  {
    path: "/",
    element: <MainLayout />,
    children: [
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
    ]
  },
])
