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
	gin.SetMode(gin.ReleaseMode)
	app = gin.New()
	app.Use(gin.Recovery())

	// Endpoint testing / health check
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Bioskop API Serverless is running",
		})
	})

	db := config.InitDB()
	if db == nil {
		app.GET("/bioskop", func(c *gin.Context) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Gagal terhubung ke database. Pastikan POSTGRES_URL dikonfigurasi di Vercel.",
			})
		})
		return
	}

	bioskopRepo := repository.NewBioskopRepository(db)
	bioskopService := service.NewBioskopService(bioskopRepo)
	bioskopHandler := handler.NewBioskopHandler(bioskopService)

	app.POST("/bioskop", bioskopHandler.CreateBioskop)
	app.GET("/bioskop", bioskopHandler.GetAllBioskop)
	app.GET("/bioskop/:id", bioskopHandler.GetBioskopByID)
	app.PUT("/bioskop/:id", bioskopHandler.UpdateBioskop)
	app.DELETE("/bioskop/:id", bioskopHandler.DeleteBioskop)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
