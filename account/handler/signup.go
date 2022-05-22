package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xblyyds/memrizr/account/model"
	"github.com/xblyyds/memrizr/account/model/apperrors"
	"log"
	"net/http"
)

// 登录请求要是私有的,就小写
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

func (h *Handler) Signup(c *gin.Context) {

	var req signupReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)

	// 注册失败
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// 注册成功
	tokens, err := h.TokenService.NewPairFromUser(c, u, "")

	if err != nil {
		log.Printf("创建token令牌失败: %v\n", err.Error())

		// 返回错误状态码及错误信息
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
