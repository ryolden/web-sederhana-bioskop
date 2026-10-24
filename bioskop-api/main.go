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

	// Endpoints CRUD Bioskop
	router.POST("/bioskop", bioskopHandler.CreateBioskop)
	router.GET("/bioskop", bioskopHandler.GetAllBioskop)
	router.GET("/bioskop/:id", bioskopHandler.GetBioskopByID)
	router.PUT("/bioskop/:id", bioskopHandler.UpdateBioskop)
	router.DELETE("/bioskop/:id", bioskopHandler.DeleteBioskop)

	router.Run(":8081")
}
