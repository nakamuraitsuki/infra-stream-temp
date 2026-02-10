package manager

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *VideoManagementHandler) UploadSource(c echo.Context) error {
	ctx := c.Request().Context()

	videoIDStr := c.Param("id")
	videoID, err := uuid.Parse(videoIDStr)
	if err != nil {
		return echo.ErrBadRequest
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get uploaded file")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to open uploaded file")
	}
	defer src.Close()

	if err := h.manageUsecase.UploadSource(ctx, videoID, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to upload video source")
	}

	return c.NoContent(http.StatusNoContent)
}
