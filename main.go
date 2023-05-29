package main

import (
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
)

func main() {
	// 目标地址
	targetURL, _ := url.Parse("http://your-new-domain.com")

	// 创建反向代理
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r := gin.Default()
	r.Any("/*any", func(c *gin.Context) {
		// 将gin的Request和ResponseWriter对象转给httputil的ReverseProxy
		proxy.ServeHTTP(c.Writer, c.Request)
	})
	r.Run()
}
