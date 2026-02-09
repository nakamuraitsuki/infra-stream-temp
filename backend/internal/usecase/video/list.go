package video

import (
	"context"

	video_domain "example.com/m/internal/domain/video"
	video_value "example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

// VideoSearchQuery defines the parameters for searching videos.
// example: limit, offset, sort order, etc.
type VideoSearchQuery struct {
	Limit int
	// add more fields as needed
}

func (uc *VideoUseCase) ListMine(
	ctx context.Context,
	ownerID uuid.UUID,
	query VideoSearchQuery,
) ([]*video_domain.Video, error) {

	cond := video_domain.ListCondition{
		OwnerID: &ownerID,
		Limit:   query.Limit,
	}

	return uc.videoRepo.FindByCondition(ctx, cond)
}

func (uc *VideoUseCase) ListPublic(
	ctx context.Context,
	query VideoSearchQuery,
) ([]*video_domain.Video, error) {

	visibility := video_value.VisibilityPublic
	status := video_value.StatusReady

	cond := video_domain.ListCondition{
		Visibility: &visibility,
		Status:     &status,
		Limit:      query.Limit,
	}

	return uc.videoRepo.FindByCondition(ctx, cond)
}

func (uc *VideoUseCase) SearchByTag(
	ctx context.Context,
	tag video_value.Tag,
	query VideoSearchQuery,
) ([]*video_domain.Video, error) {

	visibility := video_value.VisibilityPublic
	status := video_value.StatusReady

	cond := video_domain.ListCondition{
		Tag:        &tag,
		Visibility: &visibility,
		Status:     &status,
		Limit:      query.Limit,
	}

	return uc.videoRepo.FindByCondition(ctx, cond)
}
