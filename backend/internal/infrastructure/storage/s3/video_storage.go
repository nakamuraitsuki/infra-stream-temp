package s3

import (
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
