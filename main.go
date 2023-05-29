package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	// 目标服务器的 URL
	target, _ := url.Parse("https://your-domain.com")

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 创建一个 Gin 实例
	r := gin.Default()

	r.Any("/*any", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	// 运行在 8080 端口
	r.Run(":8089")
}
