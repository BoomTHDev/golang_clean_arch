package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/BoomTHDev/golang_clean_arch/databases"
)

type sessionRepositoryRedis struct {
	redis databases.RedisClient
}

func NewSessionRepositoryRedis(redis databases.RedisClient) SessionRepository {
	return &sessionRepositoryRedis{redis: redis}
}

func (r *sessionRepositoryRedis) SetSession(ctx context.Context, sessionId string, session UserSession, expiration time.Duration) error {
	key := "session-id:" + sessionId
	data, err := json.Marshal(session)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, data, expiration)
}

func (r *sessionRepositoryRedis) GetSession(ctx context.Context, sessionId string) (*UserSession, error) {
	key := "session-id:" + sessionId
	val, err := r.redis.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	session := UserSession{}
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, err
	}
	return &session, nil
}
