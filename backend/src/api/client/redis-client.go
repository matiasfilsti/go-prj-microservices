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
	SessionToken string `redis:"sessionToken"`
	CsrfToken    string `redis:"csrfToken"`
}

func (rds *ClientRedis) Get(ctx context.Context, key string) (UserSession, error) {
	var user UserSession
	if err := rds.client.HGetAll(ctx, key).Scan(&user); err != nil {
		return user, NewRedisUserNotFoundError("error getting user, user not found")
	}
	return user, nil
}

func (rds *ClientRedis) Set(ctx context.Context, key string, sessionToken string, csrfToken string) error {
	command := rds.client.TxPipeline()
	fmt.Println("Session Token:", sessionToken)
	fmt.Println("CSRF Token:", csrfToken)
	fmt.Println("User:", key)

	command.HSet(ctx, key, UserSession{sessionToken, csrfToken})
	command.Expire(ctx, key, time.Duration(config.RedisTTL)*time.Second)

	cmd, err := command.Exec(ctx)
	if err != nil {
		fmt.Println(cmd, err)
		return NewRedisUserSaveError("error saving user session tokens")
	}

	return nil
}
