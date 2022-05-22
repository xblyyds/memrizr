package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xblyyds/memrizr/account/handler"
	"github.com/xblyyds/memrizr/account/repository"
	"github.com/xblyyds/memrizr/account/service"
	"io/ioutil"
	"log"
	"os"
)

func inject(d *dataSources) (*gin.Engine, error) {
	log.Printf("注入数据源")

	userRepository := repository.NewUserRepository(d.DB)

	userService := service.NewUserService(&service.USConfig{
		UserRepository: userRepository,
	})

	privKeyFile := os.Getenv("PRIV_KEY_FILE")
	priv, err := ioutil.ReadFile(privKeyFile)

	if err != nil {
		return nil, fmt.Errorf("不能读取私钥pem文件: %w", err)
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		return nil, fmt.Errorf("不能解析私钥: %w", err)
	}

	pubKeyFile := os.Getenv("PUB_KEY_FILE")
	pub, err := ioutil.ReadFile(pubKeyFile)

	if err != nil {
		return nil, fmt.Errorf("不能读取公钥pem文件: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

	if err != nil {
		return nil, fmt.Errorf("不能解析公钥: %w", err)
	}

	refreshSecret := os.Getenv("REFRESH_SECRET")

	tokenService := service.NewTokenService(&service.TSConfig{
		PrivKey:       privKey,
		PubKey:        pubKey,
		RefreshSecret: refreshSecret,
	})

	router := gin.Default()

	handler.NewHandler(&handler.Config{
		R:            router,
		UserService:  userService,
		TokenService: tokenService,
	})

	return router, nil

}
