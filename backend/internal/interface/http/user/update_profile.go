package user

import (
	"example.com/m/internal/interface/http/middleware"
	"github.com/labstack/echo/v4"
)

type UpdateProfileRequest struct {
	Name *string `json:"name"`
	Bio  *string `json:"bio"`
}

func (h *Handler) UpdateProfile(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := middleware.GetUserID(c)
	if err != nil {
		return err
	}

	var req UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	if req.Name == nil && req.Bio == nil {
		return echo.ErrBadRequest
	}

	user, err := h.usecase.GetMe(ctx, userID)
	if err != nil {
		return echo.ErrInternalServerError
	}

	name := user.Name
	bio := user.Bio
	if req.Name != nil {
		name = *req.Name
	}
	if req.Bio != nil {
		bio = *req.Bio
	}

	if err := h.usecase.UpdateProfile(ctx, userID, name, bio); err != nil {
		return echo.ErrInternalServerError
	}

	return c.NoContent(204)
}
