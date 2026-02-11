package process

import "github.com/google/uuid"

type TranscodePayload struct {
	VideoID uuid.UUID `json:"video_id"`
}