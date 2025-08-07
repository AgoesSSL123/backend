package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/agus/my-hospital-app/config"
	"github.com/agus/my-hospital-app/models"
	"github.com/gin-gonic/gin"
)

// ➕ Tambah data pasien baru
func CreatePasien(c *gin.Context) {
	var p models.Pasien

	// Bind JSON dari frontend ke struct Pasien
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println("❌ Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("📥 Data masuk:", p)

	// Validasi wajib diisi
	if p.Nama == "" || p.Telepon == "" || p.Gender == "" {
		log.Println("⚠️ Nama, No. Telepon, dan Jenis Kelamin wajib diisi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama, No. Telepon, dan Jenis Kelamin wajib diisi"})
		return
	}

	// Query insert ke database
	err := config.DB.QueryRow(`
		INSERT INTO registrasi_pasien 
		(nama_lengkap, nik, alamat, no_telepon, bpjs, jenis_kelamin, keluhan, tanggal_registrasi)
		VALUES ($1,$2,$3,$4,$5,$6,$7,CURRENT_DATE) RETURNING id`,
		p.Nama, p.NIK, p.Alamat, p.Telepon, p.BPJS, p.Gender, p.Keluhan,
	).Scan(&p.ID)

	if err != nil {
		log.Println("❌ Error inserting pasien:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data pasien"})
		return
	}

	// Respons ke frontend
	c.JSON(http.StatusCreated, gin.H{
		"id":            p.ID,
		"nama_lengkap":  p.Nama,
		"no_telepon":    p.Telepon,
		"jenis_kelamin": p.Gender,
		"message":       "Pasien berhasil diregistrasi",
	})
}

// 👉 Ambil semua pasien
func GetAllPasien(c *gin.Context) {
	rows, err := config.DB.Query(`SELECT id, nama_lengkap, nik, alamat, no_telepon, bpjs, jenis_kelamin, keluhan FROM registrasi_pasien`)
	if err != nil {
		log.Println("❌ Error mengambil data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pasien"})
		return
	}
	defer rows.Close()

	var data []models.Pasien
	for rows.Next() {
		var p models.Pasien
		if err := rows.Scan(&p.ID, &p.Nama, &p.NIK, &p.Alamat, &p.Telepon, &p.BPJS, &p.Gender, &p.Keluhan); err != nil {
			log.Println("❌ Error scan row:", err)
			continue
		}
		data = append(data, p)
	}

	c.JSON(http.StatusOK, data)
}

// 🔍 Ambil data pasien berdasarkan ID
func GetPasienByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var p models.Pasien
	err = config.DB.QueryRow(`SELECT id, nama_lengkap, nik, alamat, no_telepon, bpjs, jenis_kelamin, keluhan FROM registrasi_pasien WHERE id = $1`, id).
		Scan(&p.ID, &p.Nama, &p.NIK, &p.Alamat, &p.Telepon, &p.BPJS, &p.Gender, &p.Keluhan)

	if err != nil {
		log.Println("❌ Error ambil data by ID:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Data pasien tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, p)
}

// ✏️ Update data pasien
func UpdatePasien(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var p models.Pasien
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = config.DB.Exec(`UPDATE registrasi_pasien SET nama_lengkap=$1, nik=$2, alamat=$3, no_telepon=$4, bpjs=$5, jenis_kelamin=$6, keluhan=$7 WHERE id=$8`,
		p.Nama, p.NIK, p.Alamat, p.Telepon, p.BPJS, p.Gender, p.Keluhan, id)

	if err != nil {
		log.Println("❌ Error update pasien:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update data pasien"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pasien berhasil diperbarui"})
}

// ❌ Hapus data pasien
func DeletePasien(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	_, err = config.DB.Exec(`DELETE FROM registrasi_pasien WHERE id=$1`, id)
	if err != nil {
		log.Println("❌ Error hapus pasien:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data pasien"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pasien berhasil dihapus"})
}