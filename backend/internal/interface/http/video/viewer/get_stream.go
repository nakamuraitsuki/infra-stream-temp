package viewer

import (
	"net/http"
	"time"

	"example.com/m/internal/usecase/video/view"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *VideoViewingHandler) GetVideoStream(c echo.Context) error {
	ctx := c.Request().Context()

	videIDStr := c.Param("id")
	videoID, err := uuid.Parse(videIDStr)
	if err != nil {
		return echo.ErrBadRequest
	}

	stream, mimeType, err := h.usecase.GetVideoStream(ctx, videoID)
	if err != nil {
		switch err {
		case view.ErrVideoNotReady:
			return echo.NewHTTPError(http.StatusConflict, "video not ready")
		case view.ErrVideoForbidden:
			return echo.NewHTTPError(http.StatusForbidden, "video is not public")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to get video stream")
		}
	}
	defer func() {
		// 形キャストしてCloseメソッドがあれば呼び出す
		if closer, ok := stream.(interface{ Close() error }); ok {
			_ = closer.Close()
		}
	}()

	req := c.Request()
	res := c.Response()

	res.Header().Set(echo.HeaderContentType, mimeType)
	res.Header().Set(echo.HeaderCacheControl, "no-store")

	http.ServeContent(
		res.Writer,
		req,
		"",
		time.Time{},
		stream,
	)

	return nil
}
