package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"example.com/m/internal/interface/http/middleware"
)

type DummyLoginResponse struct {
	Message string    `json:"message"`
	UserID  uuid.UUID `json:"userID"`
}

func (h *Handler) DummyLogin(c echo.Context) error {
	// 仮ユーザーとして UUID ゼロ値を Cookie に書き込む
	c.SetCookie(&http.Cookie{
		Name:     middleware.DummyLoginCookieName,
		Value:    uuid.Nil.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // 開発環境では false、本番では true
	})

	return c.JSON(http.StatusOK, DummyLoginResponse{
		Message: "dummy login successful",
		UserID:  uuid.Nil,
	})
}