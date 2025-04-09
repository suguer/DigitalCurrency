package router

import (
	"github.com/gin-gonic/gin"
)

// NewRouter 创建并配置路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// v1 版本
		v1 := api.Group("/v1")
		{
			// 区块链相关接口
			blockchain := v1.Group("/blockchain")
			{
				// TODO: 添加区块链相关路由
				blockchain.GET("/status", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"status": "running",
					})
				})
			}

			// 交易相关接口
			transactions := v1.Group("/transactions")
			{
				// TODO: 添加交易相关路由
				transactions.GET("/pending", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"transactions": []interface{}{},
					})
				})
			}
		}
	}

	return r
}
