package manager

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UploadSourceRequest struct {
	FileSize int64 `json:"file_size"`
}

type UploadSourceResponse struct {
	UploadID string   `json:"upload_id"`
	URLs     []string `json:"urls"`
	PartSize int64    `json:"part_size"`
	Key      string   `json:"key"`
}

func (h *VideoManagementHandler) PrepareUpload(c echo.Context) error {
	ctx := c.Request().Context()

	videoIDStr := c.Param("id")
	videoID, err := uuid.Parse(videoIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid video ID: "+err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get uploaded file: "+err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to open uploaded file: "+err.Error())
	}
	defer src.Close()

	result, err := h.manageUsecase.PrepareUploadSession(ctx, videoID, file.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to prepare upload session: "+err.Error())
	}

	res := UploadSourceResponse{
		UploadID: result.UploadID,
		URLs:     result.URLs,
		PartSize: result.PartSize,
		Key:      result.Key,
	}

	return c.JSON(http.StatusCreated, res)
}
