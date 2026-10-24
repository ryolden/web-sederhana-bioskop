package handler

import (
	"bioskop-api/entity"
	"bioskop-api/service"
	"net/http"
	"strconv"

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
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Field 'nama' dan 'lokasi' wajib diisi",
		})
		return
	}

	result, err := h.service.CreateBioskop(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Bioskop berhasil ditambahkan",
		"data":    result,
	})
}

func (h *BioskopHandler) GetAllBioskop(c *gin.Context) {
	result, err := h.service.GetAllBioskop(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Berhasil mengambil daftar bioskop",
		"data":    result,
	})
}

func (h *BioskopHandler) GetBioskopByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID bioskop harus berupa angka integer",
		})
		return
	}

	result, err := h.service.GetBioskopByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Berhasil mengambil detail bioskop",
		"data":    result,
	})
}

func (h *BioskopHandler) UpdateBioskop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID bioskop harus berupa angka integer",
		})
		return
	}

	var input entity.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format JSON tidak valid",
		})
		return
	}

	result, err := h.service.UpdateBioskop(c.Request.Context(), id, &input)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "bioskop tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data bioskop berhasil diperbarui",
		"data":    result,
	})
}

func (h *BioskopHandler) DeleteBioskop(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID bioskop harus berupa angka integer",
		})
		return
	}

	err = h.service.DeleteBioskop(c.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "bioskop tidak ditemukan" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Bioskop berhasil dihapus",
	})
}
