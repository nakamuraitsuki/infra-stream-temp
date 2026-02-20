package s3

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func NewClient(ctx context.Context, cfg Config) (*s3.Client, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				"", // session token（通常不要）
			),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
		}
		o.UsePathStyle = cfg.UsePathStyle
	})

	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(cfg.BucketName),
	})
	if err != nil {
		var ownedErr *types.BucketAlreadyOwnedByYou
		var existsErr *types.BucketAlreadyExists
		if !errors.As(err, &ownedErr) && !errors.As(err, &existsErr) {
			return nil, err
		}
	}

	return client, nil
}
