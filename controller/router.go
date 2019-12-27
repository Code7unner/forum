package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"simple-forum/controller/routes"
)

type Controller struct {
	*gin.Engine
	db *sql.DB
}

func InitController(db *sql.DB) *Controller {
	return &Controller{gin.New(), db}
}

func (c *Controller) InitMiddleware() {
	c.Use(gin.Logger())
	c.Use(c.CORSMiddleware())
}

func (c *Controller) RunServer(address string) error {
	if address == "" {
		return errors.New("Error port address, please change your PORT in .env file")
	}

	if err := c.Run(":" + address); err != nil {
		return err
	}

	return nil
}

func (c *Controller) CORSMiddleware() gin.HandlerFunc {
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

func (c *Controller) Routes() {
	c.GET("/api/forum/categories", routes.GetCategories)
}