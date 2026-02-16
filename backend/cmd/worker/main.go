package main

import (
	"example.com/m/internal/infrastructure/persistence/postgres"
	"example.com/m/internal/infrastructure/persistence/postgres/outbox"
	"example.com/m/internal/infrastructure/queue/redis"
	"example.com/m/internal/infrastructure/storage/s3"
)

func main() {
	// NOTE: Worker Queue は、Keyを切り替えることで複数用途に利用可能
	//       よって、同じ型を複数作りづらい Wire は導入しない

	jobQueueCfg := redis.QueueConfig{ Key: "job_queue" }
	redisCfg := redis.NewRedisConfig()
	postgresCfg := postgres.NewPostgresConfig()
	s3Cfg := s3.NewS3Config()

	redisClient, err := redis.NewClient(redisCfg)
	if err != nil {
		panic(err)
	}
	postgresClient, err := postgres.NewClient(postgresCfg)
	if err != nil {
		panic(err)
	}
	s3Client, err := s3.NewS3Client(s3Cfg)
	if err != nil {
		panic(err)
	}

	outboxRepo := outbox.NewPostgresOutboxRepository(postgresClient)
	videoStorage := s3.NewVideoStorage()

	jobQueue := redis.NewQueue(redisClient, jobQueueCfg.Key)
	jobConsumer := redis.NewConsumer(redisClient, jobQueue, jobQueueCfg.Key)
	relayWorker := 
}
