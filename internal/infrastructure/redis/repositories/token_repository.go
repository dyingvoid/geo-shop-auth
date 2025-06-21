package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"geo-shop-auth/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisTokenRepository struct {
	client *redis.Client
}

func (tp *RedisTokenRepository) Insert(ctx context.Context, token *domain.RefreshToken) error {
	b, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("error marshaling refresh token: %w", err)
	}

	key := fmt.Sprintf("refresh_token:%s", token.Value.String())
	err = tp.client.Set(ctx, key, b,
		time.Until(time.Unix(token.ExpTime, 0)),
	).Err()
	if err != nil {
		return fmt.Errorf("redis error setting refresh token: %w", err)
	}

	return nil
}

func (tp *RedisTokenRepository) FindToken(ctx context.Context, str string) (*domain.RefreshToken, error) {
	key := fmt.Sprintf("refresh_token:%s", str)
	val, err := tp.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("redis error getting refresh token: %w", err)
	}

	var token domain.RefreshToken
	err = json.Unmarshal([]byte(val), &token)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling refresh token: %w", err)
	}

	return &token, err
}

func NewRedisTokenRepository(client *redis.Client) *RedisTokenRepository {
	return &RedisTokenRepository{client: client}
}
