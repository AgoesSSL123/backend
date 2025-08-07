package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/agus/my-hospital-app/config"
	"github.com/agus/my-hospital-app/models"
	"github.com/gin-gonic/gin"
)

// ➕ Tambah data obat
func CreateObat(c *gin.Context) {
	var o models.Obat
	if err := c.ShouldBindJSON(&o); err != nil {
		log.Println("❌ Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if o.NamaObat == "" || o.NamaPenyakit == "" || o.Expired == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Field nama obat, nama penyakit, dan expired wajib diisi"})
		return
	}

	err := config.DB.QueryRow(`
		INSERT INTO obat (nama_penyakit, nama_obat, expired, keterangan)
		VALUES ($1, $2, $3, $4) RETURNING id`,
		o.NamaPenyakit, o.NamaObat, o.Expired, o.Keterangan,
	).Scan(&o.ID)

	if err != nil {
		log.Println("❌ Error insert:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusCreated, o)
}

// 📄 Ambil semua data
func GetAllObat(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, nama_penyakit, nama_obat, expired, keterangan FROM obat`)
	if err != nil {
		log.Println("❌ Error ambil data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil data"})
		return
	}
	defer rows.Close()

	var list []models.Obat
	for rows.Next() {
		var o models.Obat
		if err := rows.Scan(&o.ID, &o.NamaPenyakit, &o.NamaObat, &o.Expired, &o.Keterangan); err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}
		list = append(list, o)
	}
	c.JSON(http.StatusOK, list)
}

// ❌ Hapus data
func DeleteObat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	_, err = config.DB.Exec(`DELETE FROM obat WHERE id=$1`, id)
	if err != nil {
		log.Println("❌ Gagal hapus:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal hapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

// 🔍 Ambil data berdasarkan ID
func GetObatByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var o models.Obat
	err = config.DB.QueryRow(`
		SELECT id, nama_penyakit, nama_obat, expired, keterangan FROM obat WHERE id=$1`, id,
	).Scan(&o.ID, &o.NamaPenyakit, &o.NamaObat, &o.Expired, &o.Keterangan)

	if err != nil {
		log.Println("❌ Error ambil data:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, o)
}

// 📝 Update data obat
func UpdateObat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var o models.Obat
	if err := c.ShouldBindJSON(&o); err != nil {
		log.Println("❌ Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = config.DB.Exec(`
		UPDATE obat SET nama_penyakit=$1, nama_obat=$2, expired=$3, keterangan=$4 WHERE id=$5`,
		o.NamaPenyakit, o.NamaObat, o.Expired, o.Keterangan, id,
	)

	if err != nil {
		log.Println("❌ Gagal update:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}