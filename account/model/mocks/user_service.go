package mocks

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/xblyyds/memrizr/model"
)

// 模拟用户，便于测试

// model.UserService的一个模拟类型
type MockUserService struct {
	mock.Mock
}

// 模拟UserService的Get方法
func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {

	// ret.Get(0) 是返回结果 ret.Get(1) 是错误
	// 似乎是这样的,都需要强制转换
	ret := m.Called(ctx, uid)

	var r0 *model.User
	if ret.Get(0) != nil {
		// 强转一下成user
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
