package video

import (
	"example.com/m/internal/interface/http/middleware"
	"example.com/m/internal/interface/http/video/manager"
	"example.com/m/internal/interface/http/video/viewer"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, mh *manager.VideoManagementHandler, vh *viewer.VideoViewingHandler) {
	e.GET("/videos/:id", vh.GetByID)
	e.GET("/videos", vh.ListPublic)
	e.GET("/videos/search", vh.SearchByTag)

	videos := e.Group("/videos")
	videos.Use(middleware.DummyAuthMiddleware) // Dummyの認証ミドルウェアを使用

	videos.POST("", mh.Create)
	videos.GET("/mine", mh.ListMine)
}
