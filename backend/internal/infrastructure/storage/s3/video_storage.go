package s3

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"example.com/m/internal/domain/video"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
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

func (s *videoStorage) SaveSource(ctx context.Context, sourceKey string, data io.Reader) error {
	return s.upload(ctx, sourceKey, data, "video/mp4")
}

func (s *videoStorage) SaveStream(ctx context.Context, streamKey string, data io.Reader) error {
	// 拡張子から判断した方が良い
	return s.upload(ctx, streamKey, data, "application/x-mpegURL")
}

func (s *videoStorage) GenerateTemporaryAccessURL(ctx context.Context, streamKey string, expiresDuration time.Duration) (string, error) {
	return "", nil
}

// GetStream 一時ファイル作成　io.ReadSeekerで返す
func (s *videoStorage) GetStream(ctx context.Context, streamKey string) (io.ReadSeeker, error) {
	output, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(streamKey),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get stream from s3: %w", err)
	}
	defer output.Body.Close()

	tmpFile, err := os.CreateTemp("", "video-stream-*.tmp")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}

	if _, err := io.Copy(tmpFile, output.Body); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to copy to temp file: %w", err)
	}

	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("failed to seek temp file: %w", err)
	}

	return tmpFile, nil
}

func (s *videoStorage) DeleteSource(ctx context.Context, sourceKey string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(sourceKey),
	})
	return err
}

// streamKey は Prefix を渡し、配下のオブジェクトを全て削除する
func (s *videoStorage) DeleteStream(ctx context.Context, streamKey string) error {
	// cf. https://qiita.com/hkak03key/items/410fda8f008a21a4ca43
	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
		Prefix: aws.String(streamKey),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		if len(page.Contents) == 0 {
			continue
		}

		var objects []types.ObjectIdentifier
		for _, obj := range page.Contents {
			objects = append(
				objects,
				types.ObjectIdentifier{Key: obj.Key},
			)
		}

		_, err = s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(s.bucketName),
			Delete: &types.Delete{Objects: objects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Helper method to upload data to S3
func (s *videoStorage) upload(ctx context.Context, key string, data io.Reader, contentType string) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(key),
		Body:        data,
		ContentType: aws.String(contentType),
	})

	return err
}
