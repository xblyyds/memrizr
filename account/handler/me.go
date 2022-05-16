package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xblyyds/memrizr/model"
	"github.com/xblyyds/memrizr/model/apperrors"
	"log"
	"net/http"
)

// 用户的信息
func (h *Handler) Me(c *gin.Context) {
	// 检测 上下文context中是否有 user
	user, exists := c.Get("user")

	// 如果 没有user
	if !exists {
		log.Printf("不能从请求上下文中提取user,未知的原因: %v\n", c)
		err := apperrors.NewInternal()
		// 返回500错误
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	// 上下文中有user
	// 强转一下
	uid := user.(*model.User).UID

	// 调用service
	u, err := h.UserService.Get(c, uid)

	// 没有查询到该用户
	if err != nil {
		log.Printf("没有该用户: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())
		// 返回404错误
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	// 成功查询到了，返回该用户
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
