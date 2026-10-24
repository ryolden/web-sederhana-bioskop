package main

import (
	"bioskop-api/config"
	"bioskop-api/handler"
	"bioskop-api/repository"
	"bioskop-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	bioskopRepo := repository.NewBioskopRepository(db)
	bioskopService := service.NewBioskopService(bioskopRepo)
	bioskopHandler := handler.NewBioskopHandler(bioskopService)

	router := gin.Default()

	router.POST("/bioskop", bioskopHandler.CreateBioskop)

	router.Run(":8081")
}
