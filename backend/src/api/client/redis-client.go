package client

import (
	"backend/src/api/config"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type ClientRedis struct {
	client *redis.Client
}

func NewClientRedis(Client *redis.Client) *ClientRedis {
	return &ClientRedis{
		client: Client,
	}
}
func CreateRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHostname, config.RedisPort),
		Username: config.RedisUser,
		Password: config.RedisPassword,
	})
	return rdb
}

type UserSession struct {
	sessionToken string
	csrfToken    string
	expiration   time.Duration
}

func (rds *ClientRedis) Get(ctx context.Context, key string) (UserSession, error) {
	var user UserSession
	if err := rds.client.HGetAll(ctx, key).Scan(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (rds *ClientRedis) Set(ctx context.Context, key string, sessionToken string, csrfToken string) error {
	session := UserSession{
		sessionToken: sessionToken,
		csrfToken:    csrfToken,
		expiration:   3600 * time.Second,
	}

	if err := rds.client.HSet(ctx, key, session).Err(); err != nil {
		return err
	}

	if err := rds.client.Expire(ctx, key, session.expiration).Err(); err != nil {
		return err
	}
	return nil

}
