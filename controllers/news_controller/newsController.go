package newscontroller

import (
	"log"
	"news_db/config"
	"news_db/models"
	"news_db/utils"

	"github.com/gin-gonic/gin"
)

func GetNewsById(c *gin.Context) {}

func GetAllNews(c *gin.Context) {
	var news []models.News

	err := config.DbConnect().Find(&news)
	log.Println(err)
	utils.ResponseSuccess(c, news)

}

func PostNews(c *gin.Context) {}

func PutNewsById(c *gin.Context) {}
