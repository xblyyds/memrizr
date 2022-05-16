package model

import (
	"context"
	"github.com/google/uuid"
)

// service层
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
}

// dao层
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}
