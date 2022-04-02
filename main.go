package main

import (
	"api/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.Default()

	// 初始化路由绑定
	bootstrap.SetupRoute(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
