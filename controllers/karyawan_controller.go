package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/agus/my-hospital-app/config"
	"github.com/agus/my-hospital-app/models"
	"github.com/gin-gonic/gin"
)

func GetAllKaryawan(c *gin.Context) {
	rows, err := config.DB.Query("SELECT * FROM karyawan")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Karyawan
	for rows.Next() {
		var k models.Karyawan
		err := rows.Scan(&k.ID, &k.Nama, &k.NIK, &k.Whatsapp, &k.Gender, &k.Alamat, &k.Gaji, &k.TanggalBergabung)
		if err != nil {
			log.Println(err)
			continue
		}
		list = append(list, k)
	}

	c.JSON(http.StatusOK, list)
}

func CreateKaryawan(c *gin.Context) {
	var k models.Karyawan
	if err := c.ShouldBindJSON(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO karyawan (nama, nik, whatsapp, gender, alamat, gaji, tanggal_bergabung)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	err := config.DB.QueryRow(query, k.Nama, k.NIK, k.Whatsapp, k.Gender, k.Alamat, k.Gaji, k.TanggalBergabung).Scan(&k.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, k)
}

func UpdateKaryawan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var k models.Karyawan
	if err := c.ShouldBindJSON(&k); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE karyawan SET nama=$1, nik=$2, whatsapp=$3, gender=$4, alamat=$5, gaji=$6, tanggal_bergabung=$7 WHERE id=$8`
	_, err = config.DB.Exec(query, k.Nama, k.NIK, k.Whatsapp, k.Gender, k.Alamat, k.Gaji, k.TanggalBergabung, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data karyawan berhasil diupdate"})
}

func DeleteKaryawan(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	_, err = config.DB.Exec("DELETE FROM karyawan WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data karyawan berhasil dihapus"})
}
