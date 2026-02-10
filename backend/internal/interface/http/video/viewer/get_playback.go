package viewer

import (
	"net/http"

	"example.com/m/internal/usecase/video/view"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetPlaybackInfoResponse struct {
	PlaybackURL string `json:"playback_url"`
	MIMEType    string `json:"mime_type"`
}

func (h *VideoViewingHandler) GetPlaybackInfo(c echo.Context) error {
	ctx := c.Request().Context()

	pvideoIDStr := c.Param("video_id")
	videoID, err := uuid.Parse(pvideoIDStr)
	if err != nil {
		return echo.ErrBadRequest
	}

	info, err := h.usecase.GetPlaybackInfo(ctx, videoID)
	if err != nil {
		switch err {
			case view.ErrVideoNotReady:
				return echo.NewHTTPError(http.StatusConflict, "video not ready")
			case view.ErrVideoForbidden:
				return echo.NewHTTPError(http.StatusForbidden, "video is not public")
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to get playback info")
		}
	}

	// クライアントはこのURLにアクセスして動画ストリームを取得する
	playbackURL := "/api/videos/" + videoID.String() + "/stream"

	resp := GetPlaybackInfoResponse{
		PlaybackURL: playbackURL,
		MIMEType:    info.MIMEType,
	}

	return c.JSON(http.StatusOK, resp)
}
