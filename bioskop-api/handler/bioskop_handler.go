package handler

import (
	"bioskop-api/entity"
	"bioskop-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BioskopHandler struct {
	service service.BioskopService
}

func NewBioskopHandler(service service.BioskopService) *BioskopHandler {
	return &BioskopHandler{service: service}
}

func (h *BioskopHandler) CreateBioskop(c *gin.Context) {
	var input entity.Bioskop

	// Binding JSON request ke struct entity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Field 'nama' dan 'lokasi' wajib diisi",
		})
		return
	}

	// Panggil service layer
	result, err := h.service.CreateBioskop(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Response sukses (201 Created)
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Bioskop berhasil ditambahkan",
		"data":    result,
	})
}
