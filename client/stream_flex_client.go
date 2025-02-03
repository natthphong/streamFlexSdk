package client

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/natthphong/streamFlexSdk/kafka"
	"net/http"
)

type StreamFlexClient struct {
	Ctx           context.Context
	DB            *pgxpool.Pool
	RedisClient   *redis.UniversalClient
	HTTPClient    *http.Client
	S3Client      *s3.S3
	KafkaProducer *kafka.SendMessageSyncWithTopicFunc
	details       map[string]interface{}
	Payload       []byte
}
type StreamFlexScript interface {
	Process(client *StreamFlexClient) error
}

func NewStreamFlexClient(
	ctx context.Context,
	db *pgxpool.Pool,
	redisClient *redis.UniversalClient,
	httpClient *http.Client,
	s3Client *s3.S3,
	kafkaProducer *kafka.SendMessageSyncWithTopicFunc,
	payload []byte,
) *StreamFlexClient {
	return &StreamFlexClient{
		Ctx:           ctx,
		DB:            db,
		RedisClient:   redisClient,
		HTTPClient:    httpClient,
		S3Client:      s3Client,
		KafkaProducer: kafkaProducer,
		Payload:       payload,
	}
}
