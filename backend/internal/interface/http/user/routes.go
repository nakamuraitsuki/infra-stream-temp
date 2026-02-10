package user

import (
	"example.com/m/internal/interface/http/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/users", h.Register)
	e.POST("/users/login", h.DummyLogin) // 認証不要な仮ログインエンドポイント
	e.POST("/users/logout", h.Logout)

	auth := e.Group("/users")
	auth.Use(middleware.DummyAuthMiddleware)
	
	auth.GET("/me", h.GetMe)
	auth.PUT("/me/profile", h.UpdateProfile)
	auth.PUT("/me/icon", h.UpdateIcon)
}
