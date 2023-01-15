package main

import (
	"log"
	"news_db/config"
	"news_db/controllers/news_controller"
	

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	db := config.DbConnect()
	defer db.Close()

	router.GET("/api/news", newscontroller.GetAllNews)
	router.GET("/api/news/:id", newscontroller.GetNewsById)
	router.POST("/api/news", newscontroller.PostNews)
	router.PUT("/api/news/:id", newscontroller.PutNewsById)


	err := router.Run(":" + "3000")
	if err != nil {
		log.Fatal(err)
	}
}
