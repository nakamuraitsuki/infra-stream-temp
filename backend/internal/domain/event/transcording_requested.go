package event

import (
	"time"

	"github.com/google/uuid"
)

type TranscodingRequested struct {
	id         uuid.UUID
	VideoID    uuid.UUID
	occurredAt time.Time
}

func NewTranscodingRequested(videoID uuid.UUID) Event {
	return &TranscodingRequested{
		id:         uuid.New(),
		VideoID:    videoID,
		occurredAt: time.Now(),
	}
}

func (e *TranscodingRequested) ID() uuid.UUID {
	return e.id
}

func (e *TranscodingRequested) EventType() string {
	return "video.transcoding_requested"
}

func (e *TranscodingRequested) OccurredAt() time.Time {
	return e.occurredAt
}
