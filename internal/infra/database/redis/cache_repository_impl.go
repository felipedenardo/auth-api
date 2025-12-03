package redis

import (
	"context"
	"github.com/felipedenardo/chameleon-auth-api/internal/domain/auth"
	"github.com/redis/go-redis/v9"
	"time"
)

type cacheRepository struct {
	client *redis.Client
}

func NewCacheRepository(client *redis.Client) auth.ICacheRepository {
	return &cacheRepository{client: client}
}

func (r *cacheRepository) BlacklistToken(ctx context.Context, jti string, ttl time.Duration) error {
	return r.client.Set(ctx, jti, "1", ttl).Err()
}

func (r *cacheRepository) IsTokenBlacklisted(ctx context.Context, jti string) (bool, error) {
	cmd := r.client.Exists(ctx, jti)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	return cmd.Val() == 1, nil
}
