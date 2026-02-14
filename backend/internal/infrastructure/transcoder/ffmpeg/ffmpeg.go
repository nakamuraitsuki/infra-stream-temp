package ffmpeg

import (
	"context"

	"example.com/m/internal/domain/video"
)

type FfmpegTranscoder struct {
	storage video.Storage
}

func NewFfmpegTranscoder(storage video.Storage) video.Transcoder {
	return &FfmpegTranscoder{
		storage: storage,
	}
}

func (t *FfmpegTranscoder) Transcode(ctx context.Context, sourceKey string, streamKey string) error {
	return nil
}
