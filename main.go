package main

//main.go
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/endpoint"
)

func main() {
	r := gin.Default() // 返回一个默认的路由引擎

	//get请求 访问路径"/hello 执行匿名函数
	r.GET("/hello", func(ctx *gin.Context) {
		//返回一个json格式的响应
		ctx.JSON(200, gin.H{
			"message": "Hello Golang Gin",
		})
	})
	r.POST("/answer", endpoint.NewHelloController().Handle)
	//启动服务
	err := r.Run(":9090")
	if err != nil {
		fmt.Println("http serve failed, err =", err)
	} else {
		fmt.Println("http serve success")
	}
}
