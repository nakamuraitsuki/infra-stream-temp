package user

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo, h *Handler) {
	e.POST("/users", h.Register)

	auth := e.Group("/users")
	// TODO: Add authentication middleware here
	
	auth.GET("/me", h.GetMe)
	auth.PUT("/me/profile", h.UpdateProfile)
	auth.PUT("/me/icon", h.UpdateIcon)
}
