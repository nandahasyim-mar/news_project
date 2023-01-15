package utils

import (
	"net/http"
	"news_db/models"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	responData := models.Respon{
		Status:  "200",
		Message: "Success get data!!",
		Data:    data,
	}

	c.JSON(http.StatusOK, responData)
}
