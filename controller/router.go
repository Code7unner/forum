package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func InitRouter() *Router {
	return &Router{gin.New()}
}

func (r *Router) InitMiddleware() {
	r.Use(gin.Logger())
	r.Use(r.CORSMiddleware())
}

func (r *Router) RunServer(address string) error {
	if address == "" {
		return errors.New("Error port address, please change your PORT in .env file")
	}

	if err := r.Run(":" + address); err != nil {
		return err
	}

	return nil
}

func (r *Router) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (r *Router) Routes() {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}