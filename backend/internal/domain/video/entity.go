package video

import (
	"time"

	"example.com/m/internal/domain/video/value"
)

type Video struct {
	id            string
	ownerID       string
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
