package viewer

import (
	"example.com/m/internal/usecase/video/query"
	"github.com/labstack/echo/v4"
)

type SearchByTagRequest struct {
	Tag   string `query:"tag"`
	Limit int    `query:"limit"`
}

type SearchByTagResponse struct {
	Items []SearchByTagResponseItem `json:"items"`
}

type SearchByTagResponseItem struct {
	ID          string   `json:"id"`
	OwnerID     string   `json:"owner_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"created_at"`
}

func (h *VideoViewingHandler) SearchByTag(c echo.Context) error {
	ctx := c.Request().Context()

	var req SearchByTagRequest
	if err := echo.QueryParamsBinder(c).
		String("tag", &req.Tag).
		Int("limit", &req.Limit).
		BindError(); err != nil {
		return echo.ErrBadRequest
	}

	result, err := h.usecase.SearchByTag(ctx, req.Tag, query.VideoSearchQuery{
		Limit: req.Limit,
	})
	if err != nil {
		return echo.ErrInternalServerError
	}

	items := make([]SearchByTagResponseItem, len(result.Results))
	for i, item := range result.Results {
		items[i] = SearchByTagResponseItem{
			ID:          item.ID.String(),
			OwnerID:     item.OwnerID.String(),
			Title:       item.Title,
			Description: item.Description,
			Tags:        item.Tags,
			CreatedAt:   item.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	resp := SearchByTagResponse{
		Items: items,
	}

	return c.JSON(200, resp)
}
