package model

import (
	"context"
	"github.com/google/uuid"
)

// service层
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)

	Signup(ctx context.Context, u *User) error
}

// 产生token
type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

// dao层
type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	Create(ctx context.Context, u *User) error
}
