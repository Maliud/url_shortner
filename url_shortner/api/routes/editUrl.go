package routes

import (
	"net/http"
	"time"

	"github.com/Maliud/url_shortner/api/database"
	"github.com/Maliud/url_shortner/api/models"
	"github.com/gin-gonic/gin"
)

func EditURl(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Parse JSON",
		})
	}

	r := database.CreateClient(0)
	defer r.Close()

	// check if the ShortID exists in the DB or not
	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error" : "ShortID doesnt exists",
		})
	} 

	// update the content of the URL, expirt time with the shortID

	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry * 3600 * time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Unable to update the shorted content",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "The Content Has Been Updated !!!",
	})


}
