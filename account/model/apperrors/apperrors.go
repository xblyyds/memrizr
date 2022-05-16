package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

// c语言的#define
type Type string

const (
	Authorization   Type = "AUTHORIZATION"   // 没有权限
	BadRequest      Type = "BADREQUEST"      // 输入错误
	Conflict        Type = "CONFLICT"        // 已经存在（创建用户）
	Internal        Type = "INTERNAL"        // 500错误
	Notfound        Type = "NOTFOUND"        // 404未找到
	PayloadTooLarge Type = "PAYLOADTOOLARGE" // 上传文件超过限制
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error的方法 返回 错误消息message
func (e *Error) Error() string {
	return e.Message
}

// Error 的方法 根据错误类型type 返回状态码
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized // 401
	case BadRequest:
		return http.StatusBadRequest // 401
	case Conflict:
		return http.StatusConflict // 409
	case Internal:
		return http.StatusInternalServerError // 500
	case Notfound:
		return http.StatusNotFound // 404
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge // 413
	default:
		return http.StatusInternalServerError
	}
}

// 如果error是Error类型将返回状态码
func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

// 错误工厂（构造方法）

func NewAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: reason,
	}
}

func NewConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error"),
	}
}

func NewNotFound(name string, value string) *Error {
	return &Error{
		Type:    Notfound,
		Message: fmt.Sprintf(""),
	}
}

func NewPayloadTooLarge(maxBodySize int64, contendLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("最大上传大小: %v,实际下载大小: %v", maxBodySize, contendLength),
	}
}
