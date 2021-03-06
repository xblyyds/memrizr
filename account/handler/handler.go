package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xblyyds/memrizr/account/model"
	"net/http"
	"os"
)

// 有点类似java的controller

type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
}

func NewHandler(c *Config) {
	// 创建一个handler
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))

	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

}

// 登录
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "登录成功",
	})
}

// 注销
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "注销成功",
	})
}

// tokens
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "这是tokens",
	})
}

// 头像
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "这是头像",
	})
}

// 删除头像
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "删除头像成功",
	})
}

// 更新个人信息
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "个人信息更新成功",
	})
}
