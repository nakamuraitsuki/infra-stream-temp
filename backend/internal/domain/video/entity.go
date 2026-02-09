package video

import (
	"errors"
	"time"

	"example.com/m/internal/domain/video/value"
	"github.com/google/uuid"
)

type Video struct {
	id            uuid.UUID
	ownerID       uuid.UUID
	sourceKey     string // 元動画の保存先キー
	streamKey     string // 配信・再生用動画の保存先キー
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

func (v *Video) Status() value.Status {
	return v.status
}

func (v *Video) MarkUploaded(sourceKey string) error {
	if v.status != value.StatusInitial {
		return errors.New("video is not in initial state")
	}
	v.sourceKey = sourceKey
	v.status = value.StatusUploaded
	return nil
}

func (v *Video) StartTranscoding() error {
	if v.status != value.StatusUploaded {
		return errors.New("video is not transcodable")
	}
	v.status = value.StatusProcessing
	return nil
}

func (v *Video) MarkTranscodeSucceeded() {
	v.status = value.StatusReady
	v.failureReason = nil
	v.retryCount = 0
}

func (v *Video) MarkTranscodeFailed(reason value.FailureReason) {
	v.status = value.StatusFailed
	v.failureReason = &reason
	v.retryCount++
}
