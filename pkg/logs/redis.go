package logs

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisPublish struct {
	client  *redis.Client
	channel string
}

func NewRedisPublish(client *redis.Client, channel string) *RedisPublish {
	return &RedisPublish{
		client:  client,
		channel: channel,
	}
}

// Log logs a message to the RedisPublish channel.
func (r *RedisPublish) Log(s string) {
	err := r.client.Publish(r.client.Context(), r.channel, s).Err()
	if err != nil {
		log.Print(fmt.Errorf("publishing to RedisPublish: %w", err))
	}
}
