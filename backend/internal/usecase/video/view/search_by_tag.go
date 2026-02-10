package view

import (
	"context"
	"time"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"example.com/m/internal/usecase/video/query"
	"github.com/google/uuid"
)

type GetByTagsResults struct {
	Results []*GetByTagsResult
}

type GetByTagsResult struct {
	ID          uuid.UUID
	Title       string
	Description string
	OwnerID     uuid.UUID
	CreatedAt   time.Time
}

func (uc *VideoViewingUseCase) SearchByTag(
	ctx context.Context,
	tag video_value.Tag,
	query query.VideoSearchQuery,
) (*GetByTagsResults, error) {

	visibility := video_value.VisibilityPublic
	status := video_value.StatusReady

	cond := video_domain.ListCondition{
		Tag:        &tag,
		Visibility: &visibility,
		Status:     &status,
		Limit:      query.Limit,
	}

	videos, err := uc.VideoRepo.FindByCondition(ctx, cond)
	if err != nil {
		return nil, err
	}

	results := make([]*GetByTagsResult, len(videos))
	for i, video := range videos {
		results[i] = &GetByTagsResult{
			ID:          video.ID(),
			Title:       video.Title(),
			Description: video.Description(),
			OwnerID:     video.OwnerID(),
			CreatedAt:   video.CreatedAt(),
		}
	}
	
	return &GetByTagsResults{
		Results: results,
	}, nil
}
