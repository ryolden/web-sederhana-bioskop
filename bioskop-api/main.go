package main

import (
	"bioskop-api/config"
	"bioskop-api/handler"
	"bioskop-api/repository"
	"bioskop-api/service"
	"log"
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
	router := gin.Default()

	// Route utama untuk cek status server
	router.GET("/", func(c *gin.Context) {
		hasEnv := os.Getenv("POSTGRES_URL") != "" || os.Getenv("DATABASE_URL") != ""
		c.JSON(http.StatusOK, gin.H{
			"status":     "ok",
			"message":    "Bioskop API Server is running",
			"has_db_env": hasEnv,
		})
	})

	db := config.InitDB()

	// Hanya inisialisasi route database jika koneksi DB berhasil
	if db != nil {
		defer db.Close()

		bioskopRepo := repository.NewBioskopRepository(db)
		bioskopService := service.NewBioskopService(bioskopRepo)
		bioskopHandler := handler.NewBioskopHandler(bioskopService)

		router.POST("/bioskop", bioskopHandler.CreateBioskop)
		router.GET("/bioskop", bioskopHandler.GetAllBioskop)
		router.GET("/bioskop/:id", bioskopHandler.GetBioskopByID)
		router.PUT("/bioskop/:id", bioskopHandler.UpdateBioskop)
		router.DELETE("/bioskop/:id", bioskopHandler.DeleteBioskop)
	} else {
		router.GET("/bioskop", func(c *gin.Context) {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Database tidak terhubung. Cek POSTGRES_URL di Vercel Environment Variables.",
			})
		})
	}

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
