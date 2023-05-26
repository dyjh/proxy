package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/endpoint", func(c *gin.Context) {
		targetURL := "https://api.openai.com/v1/chat/completions" // 请替换为你的目标API URL

		// 读取客户端发送过来的原始POST数据
		postData, _ := ioutil.ReadAll(c.Request.Body)

		// 创建新的HTTP请求
		req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(postData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// 设置请求头
		req.Header.Set("Authorization", "Bearer xxx")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("OpenAI-Organization", "xxxxx")

		// 发送HTTP请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer resp.Body.Close()

		// 读取响应数据
		responseData, _ := ioutil.ReadAll(resp.Body)

		// 输出响应数据
		c.String(http.StatusOK, string(responseData))
	})

	router.Run(":8089") // 默认在0.0.0.0:8080启动服务
}
