package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello,golang")
	r := gin.Default()
	fmt.Println(r)
}
