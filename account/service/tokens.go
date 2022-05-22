package service

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/xblyyds/memrizr/account/model"
	"log"
	"time"
)

type IDTokenCustomClaims struct {
	User *model.User `json:"user"`
	jwt.StandardClaims
}

// 生成token
func generateIDToken(u *model.User, key *rsa.PrivateKey) (string, error) {
	unixTime := time.Now().Unix()
	tokenExp := unixTime + 60*15 // 15分钟过期时间

	claims := IDTokenCustomClaims{
		User: u,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(key)

	if err != nil {
		log.Println("未能标志令牌字符串")
		return "", err
	}

	return ss, nil
}

type RefreshToken struct {
	SS        string
	ID        string
	ExpiresIn time.Duration
}

type RefreshTokenCustomClaims struct {
	UID uuid.UUID `json:"uid"`
	jwt.StandardClaims
}

func generateRefreshToken(uid uuid.UUID, key string) (*RefreshToken, error) {
	currentTime := time.Now()
	tokenExp := currentTime.AddDate(0, 0, 3) // 3天
	tokenID, err := uuid.NewRandom()

	if err != nil {
		log.Printf("令牌ID刷新失败")
		return nil, err
	}

	claims := RefreshTokenCustomClaims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  currentTime.Unix(),
			ExpiresAt: tokenExp.Unix(),
			Id:        tokenID.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(key))

	if err != nil {
		log.Println("标识令牌字符串刷新失败")
		return nil, err
	}

	return &RefreshToken{
		SS:        ss,
		ID:        tokenID.String(),
		ExpiresIn: tokenExp.Sub(currentTime),
	}, nil
}
