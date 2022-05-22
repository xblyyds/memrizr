package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/xblyyds/memrizr/account/model"
	"github.com/xblyyds/memrizr/account/model/apperrors"
	"log"
)

type UserService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// 实现 UserService接口的方法
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
	// 密码加密
	pw, err := hashPassword(u.Password)

	if err != nil {
		log.Printf("不能注册该用户，邮箱：%v\n", u.Email)
		return apperrors.NewInternal()
	}

	// 将该用户密码更改为加密后的密码
	u.Password = pw

	if err := s.UserRepository.Create(ctx, u); err != nil {
		return err
	}

	// if we get around to adding events, we'd Publish it here
	// err := s.EventsBroker.PublishUserUpdated(u, true)

	// if err != nil {
	//  return nil, apperrors.NewInternal()
	// }

	return nil

}
