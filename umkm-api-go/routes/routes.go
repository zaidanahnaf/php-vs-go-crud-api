package routes

import (
	"umkm-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/umkm", handlers.GetUMKM)
	r.POST("/umkm", handlers.CreateUMKM)
	r.PUT("/umkm/:id", handlers.UpdateUMKM)    // Update UMKM berdasarkan ID
	r.DELETE("/umkm/:id", handlers.DeleteUMKM) // Hapus UMKM berdasarkan ID

	return r
}
