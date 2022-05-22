package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"strings"
)

// 加密密码
func hashPassword(password string) (string, error) {
	// 创建密钥
	salt := make([]byte, 32)
	// 随机产生32位密钥
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	// 将加密的密码与密钥盐拼接
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

// 应该是解密
func comparePasswords(storedPassword string, suppliedPassword string) (bool, error) {
	// 将加密的密码分成加密的密码和密钥盐
	pwsalt := strings.Split(storedPassword, ".")

	// 检测密钥盐
	salt, err := hex.DecodeString(pwsalt[1])

	// 验证失败
	if err != nil {
		return false, fmt.Errorf("不能校验用户密码")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	//
	return hex.EncodeToString(shash) == pwsalt[0], nil
}
