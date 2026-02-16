package s3

import "example.com/m/internal/infrastructure/env"

type Config struct {
	Endpoint     string
	Region       string
	BucketName   string
	AccessKey    string
	SecretKey    string
	UsePathStyle bool
	UseSSL       bool
}

func NewS3Config() Config {
	return Config{
		// NOTE: Default は MinIO のローカルを意識
		Endpoint:   env.GetString("S3_ENDPOINT", "http://localhost:9000"),
		Region:     env.GetString("S3_REGION", "us-east-1"), // MinIOでは任意だが指定が必要
		BucketName: env.GetString("S3_BUCKET", "my-video-app"),
		AccessKey:  env.GetString("S3_ACCESS_KEY", "minioadmin"),
		SecretKey:  env.GetString("S3_SECRET_KEY", "minioadmin"),

		// MinIOを使う場合、以下の2つは重要
		UsePathStyle: env.GetBool("S3_USE_PATH_STYLE", true), // MinIO は true 必須
		UseSSL:       env.GetBool("S3_USE_SSL", false),
	}
}
