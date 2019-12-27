package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @route  GET api/forum/categories
// @desc   Get all categories fields from db
// @access Private
func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all categories from db!"})
}
