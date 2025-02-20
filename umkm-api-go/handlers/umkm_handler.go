package handlers

import (
	"net/http"
	"umkm-api-go/config"
	"umkm-api-go/models"

	"github.com/gin-gonic/gin"
)

// GET: Ambil semua UMKM
func GetUMKM(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, nama_usaha, pemilik, kategori FROM umkm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var umkmList []models.UMKM
	for rows.Next() {
		var u models.UMKM
		if err := rows.Scan(&u.ID, &u.NamaUsaha, &u.Pemilik, &u.Kategori); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data"})
			return
		}
		umkmList = append(umkmList, u)
	}
	c.JSON(http.StatusOK, umkmList)
}

// POST: Tambah UMKM
func CreateUMKM(c *gin.Context) {
	var umkm models.UMKM
	if err := c.ShouldBindJSON(&umkm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	_, err := config.DB.Exec("INSERT INTO umkm (nama_usaha, pemilik, kategori) VALUES (?, ?, ?)",
		umkm.NamaUsaha, umkm.Pemilik, umkm.Kategori)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambah UMKM"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UMKM berhasil ditambahkan"})
}

// PUT: Update data UMKM berdasarkan ID
func UpdateUMKM(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari URL
	var umkm models.UMKM

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&umkm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah"})
		return
	}

	// Eksekusi query update
	result, err := config.DB.Exec("UPDATE umkm SET nama_usaha=?, pemilik=?, kategori=? WHERE id=?",
		umkm.NamaUsaha, umkm.Pemilik, umkm.Kategori, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data"})
		return
	}

	// Cek apakah ada baris yang terpengaruh
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "UMKM tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UMKM berhasil diperbarui"})
}

// DELETE: Hapus data UMKM berdasarkan ID
func DeleteUMKM(c *gin.Context) {
	id := c.Param("id")

	// Eksekusi query delete
	result, err := config.DB.Exec("DELETE FROM umkm WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	// Cek apakah ada baris yang terpengaruh
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "UMKM tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UMKM berhasil dihapus"})
}
