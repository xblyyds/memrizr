package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xblyyds/memrizr/model"
	"github.com/xblyyds/memrizr/model/apperrors"
	"github.com/xblyyds/memrizr/model/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMe(t *testing.T) {
	// 设置成测试模式
	gin.SetMode(gin.TestMode)

	// 成功
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UID:   uid,
			Email: "qq2368269411@163.com",
			Name:  "但求一败",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				UID: uid,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
	})

	// 没有成功
	t.Run("NoContextUser", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

		rr := httptest.NewRecorder()

		router := gin.Default()
		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.Equal(t, 500, rr.Code)
		mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	// 找不到用户
	t.Run("NotFound", func(t *testing.T) {
		uid, _ := uuid.NewRandom()
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("有一些错误调用链"))

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				UID: uid,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respErr := apperrors.NewNotFound("user", uid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)

		assert.Equal(t, respErr.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)
	})

}
