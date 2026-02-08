package video

import (
	"time"

	"example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type Video struct {
	id            uuid.UUID
	ownerID       uuid.UUID
	sourceKey     string
	streamKey     string
	status        value.Status
	title         string
	description   string
	tags          []value.Tag
	retryCount    int
	failureReason *value.FailureReason
	visibility    value.Visibility
	createdAt     time.Time
}

func NewVideo(
	id uuid.UUID,
	ownerID uuid.UUID,
	sourceKey string,
	streamKey string,
	status value.Status,
	title string,
	description string,
	tags []value.Tag,
	visibility value.Visibility,
	createdAt time.Time,
) *Video {
	return &Video{
		id:          id,
		ownerID:     ownerID,
		sourceKey:   sourceKey,
		streamKey:   streamKey,
		status:      status,
		title:       title,
		description: description,
		tags:        tags,
		visibility:  visibility,
		createdAt:   createdAt,
	}
}