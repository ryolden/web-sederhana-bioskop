package main

import (
	"bioskop-api/config"
	"bioskop-api/handler"
	"bioskop-api/repository"
	"bioskop-api/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Endpoint Diagnostik (Akan selalu HTTP 200 OK)
	r.GET("/", func(c *gin.Context) {
		hasEnv := os.Getenv("POSTGRES_URL") != "" || os.Getenv("DATABASE_URL") != ""
		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"message":    "Bioskop API Server is running",
			"has_db_env": hasEnv,
		})
	})

	db := config.InitDB()

	// Hanya daftarkan route DB jika db TIDAK nil (mencegah nil pointer crash)
	if db != nil {
		bioskopRepo := repository.NewBioskopRepository(db)
		bioskopService := service.NewBioskopService(bioskopRepo)
		bioskopHandler := handler.NewBioskopHandler(bioskopService)

		r.POST("/bioskop", bioskopHandler.CreateBioskop)
		r.GET("/bioskop", bioskopHandler.GetAllBioskop)
		r.GET("/bioskop/:id", bioskopHandler.GetBioskopByID)
		r.PUT("/bioskop/:id", bioskopHandler.UpdateBioskop)
		r.DELETE("/bioskop/:id", bioskopHandler.DeleteBioskop)
	} else {
		r.GET("/bioskop", func(c *gin.Context) {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Database belum terhubung. Pastikan POSTGRES_URL sudah ada di Vercel Settings > Environment Variables.",
			})
		})
	}

	r.Run(":" + port)
}
