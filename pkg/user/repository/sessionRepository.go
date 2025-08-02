package repository

import (
	"context"
	"time"
)

type UserSession struct {
	ID   string
	Role string
}

type SessionRepository interface {
	SetSession(ctx context.Context, sessionId string, session UserSession, expiration time.Duration) error
	GetSession(ctx context.Context, sessionId string) (*UserSession, error)
}
