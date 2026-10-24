package handler

import (
	"bioskop-api/config"
	"bioskop-api/handler"
	"bioskop-api/repository"
	"bioskop-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func init() {
	db := config.InitDB()

	bioskopRepo := repository.NewBioskopRepository(db)
	bioskopService := service.NewBioskopService(bioskopRepo)
	bioskopHandler := handler.NewBioskopHandler(bioskopService)

	app = gin.Default()

	app.POST("/bioskop", bioskopHandler.CreateBioskop)
	app.GET("/bioskop", bioskopHandler.GetAllBioskop)
	app.GET("/bioskop/:id", bioskopHandler.GetBioskopByID)
	app.PUT("/bioskop/:id", bioskopHandler.UpdateBioskop)
	app.DELETE("/bioskop/:id", bioskopHandler.DeleteBioskop)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
