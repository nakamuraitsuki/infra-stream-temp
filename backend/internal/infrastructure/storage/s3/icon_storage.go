package s3

import (
	"context"

	"example.com/m/internal/domain/user"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type iconStorage struct {
	client *s3.Client
	bucketName string
}

func NewIconStorage(client *s3.Client, bucketName string) user.IconStorage {
	return &iconStorage{
		client: client,
		bucketName: bucketName,
	}
}

func (s *iconStorage) Upload(ctx context.Context, key string, data []byte) error {
	// implement
	return nil
}

func (s *iconStorage) Delete(ctx context.Context, key string) error {
	// implement
	return nil
}