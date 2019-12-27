package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	DB *sql.DB
}

func InitController(db *sql.DB) *Controller {
	return &Controller{db}
}

func InitMiddleware(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(CORSMiddleware())
}

func RunServer(address string, c *Controller) error {
	if address == "" {
		return errors.New("Error port address, please change your PORT in .env file")
	}

	r := gin.New()

	InitMiddleware(r)
	Routes(c, r)

	if err := r.Run(":" + address); err != nil {
		return err
	}

	return nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}

func Routes(c *Controller, r *gin.Engine)  {
	r.GET("/api/forum/categories", gin.WrapF(c.GetCategories))
}