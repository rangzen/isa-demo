package subscriber

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisSubscriber struct {
	client  *redis.Client
	channel string
}

func NewRedisSubscriber(client *redis.Client, channel string) *RedisSubscriber {
	return &RedisSubscriber{
		client:  client,
		channel: channel,
	}
}

// Subscribe subscribes to a Redis channel and calls the callback function
// for each message received.
func (s *RedisSubscriber) Subscribe(callback func(msg string)) error {
	pubsub := s.client.Subscribe(s.client.Context(), s.channel)
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			log.Fatalf("closing Redis subscriber: %v", err)
		}
	}(pubsub)

	ch := pubsub.Channel()
	for msg := range ch {
		callback(fmt.Sprintf("redis:%s:%s", s.channel, msg.Payload))
	}

	return nil
}
