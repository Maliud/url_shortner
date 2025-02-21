package routes

import (
	"net/http"

	"github.com/Maliud/url_shortner/api/database"
	"github.com/gin-gonic/gin"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	err := r.Del(database.Ctx, shortID).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "enable to Delete shortened Link",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Shortened URL Deleted Successfully",
	})
}