package s3

import (
	"context"
	"io"

	"example.com/m/internal/domain/video"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type videoStorage struct {
	client     *s3.Client
	bucketName string
}

func NewVideoStorage(client *s3.Client, bucketName string) video.Storage {
	return &videoStorage{
		client:     client,
		bucketName: bucketName,
	}
}

func (s *videoStorage) SaveSource(ctx context.Context, key string, data io.Reader) error {
	// Implementation to save the source video to S3
	return nil
}

func (s *videoStorage) SaveStream(ctx context.Context, key string, data io.Reader) error {
	// Implementation to save the transcoded video to S3
	return nil
}

func (s *videoStorage) GetStream(ctx context.Context, key string) (io.ReadSeeker, error) {
	// Implementation to retrieve the transcoded video from S3
	return nil, nil
}

func (s *videoStorage) Delete(ctx context.Context, key string) error {
	// Implementation to delete the video from S3
	return nil
}
