package routes

import (
  "github.com/agus/my-hospital-app/controllers"
  "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
  api := r.Group("/api")
  {
    // 📋 Registrasi Pasien
    api.POST("/registrasi", controllers.CreatePasien)       // ➕ Tambah data
    api.GET("/registrasi", controllers.GetAllPasien)        // 📄 Ambil semua data
    api.GET("/registrasi/:id", controllers.GetPasienByID)   // 🔍 Ambil satu data
    api.PUT("/registrasi/:id", controllers.UpdatePasien)    // ✏️ Ubah data
    api.DELETE("/registrasi/:id", controllers.DeletePasien) // ❌ Hapus data

    // 🩺 Jadwal Dokter
    api.GET("/jadwal-dokter", controllers.GetJadwalDokter)
    api.POST("/jadwal-dokter", controllers.CreateJadwalDokter)
    api.DELETE("/jadwal-dokter/:id", controllers.DeleteJadwalDokter)


    // 🔑 Registrasi pengguna (opsional, jika kamu punya fitur login)
    api.POST("/register", controllers.Register)
    api.POST("/login", controllers.Login)


        // 💊 Data Obat
    api.GET("/obat", controllers.GetAllObat)         // 📄 Ambil semua obat
    api.POST("/obat", controllers.CreateObat)        // ➕ Tambah obat baru
    api.GET("/obat/:id", controllers.GetObatByID)    // 🔍 Ambil satu obat
    api.PUT("/obat/:id", controllers.UpdateObat)     // ✏️ Ubah data obat
    api.DELETE("/obat/:id", controllers.DeleteObat)  // ❌ Hapus obat

    // 👨‍💼 Data Karyawan
		api.GET("/karyawan", controllers.GetAllKaryawan)       // 📄 Ambil semua data
		api.POST("/karyawan", controllers.CreateKaryawan)      // ➕ Tambah karyawan
		api.PUT("/karyawan/:id", controllers.UpdateKaryawan)   // ✏️ Ubah karyawan
		api.DELETE("/karyawan/:id", controllers.DeleteKaryawan) // ❌ Hapus karyawan
  }
}
