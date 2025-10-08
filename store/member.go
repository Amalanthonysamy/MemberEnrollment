package store

import (
	"context"
	"time"
)

type Member struct {
	ID        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Store interface {
	CreateMember(ctx context.Context, member *Member) error
	GetMember(ctx context.Context, id string) (*Member, error)
	UpdateMember(ctx context.Context, member *Member) error
	DeleteMember(ctx context.Context, id string) error
}