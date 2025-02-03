package kafka

import (
	"go.uber.org/zap"
)

type SendMessageAsyncFunc func(event interface{})

type SendMessageSyncFunc func(log *zap.Logger, event interface{}) error
type SendMessageSyncWithTopicFunc func(log *zap.Logger, event interface{}, topic string) error
