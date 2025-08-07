package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/agus/my-hospital-app/config"
)

// 🩺 Ambil semua jadwal
func GetJadwalDokter(c *gin.Context) {
  rows, err := config.DB.Query("SELECT id, nama_dokter, spesialis, jadwal FROM jadwal_dokter")
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
    return
  }
  defer rows.Close()

  var jadwalList []gin.H
  for rows.Next() {
    var id int
    var nama, spesialis, jadwal string
    rows.Scan(&id, &nama, &spesialis, &jadwal)
    jadwalList = append(jadwalList, gin.H{
      "id":          id,
      "nama_dokter": nama,
      "spesialis":   spesialis,
      "jadwal":      jadwal,
    })
  }

  c.JSON(http.StatusOK, jadwalList)
}

// ➕ Tambah jadwal dokter
func CreateJadwalDokter(c *gin.Context) {
  var input struct {
    NamaDokter string `json:"nama_dokter"`
    Spesialis  string `json:"spesialis"`
    Jadwal     string `json:"jadwal"`
  }

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
    return
  }

  _, err := config.DB.Exec(
    "INSERT INTO jadwal_dokter (nama_dokter, spesialis, jadwal) VALUES ($1, $2, $3)",
    input.NamaDokter, input.Spesialis, input.Jadwal,
  )
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "Jadwal dokter berhasil ditambahkan"})
}

// ❌ Hapus jadwal dokter
func DeleteJadwalDokter(c *gin.Context) {
  id := c.Param("id")

  _, err := config.DB.Exec("DELETE FROM jadwal_dokter WHERE id = $1", id)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "Jadwal dokter berhasil dihapus"})
}