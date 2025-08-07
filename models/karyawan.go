package models

type Karyawan struct {
	ID              int     `json:"id"`
	Nama            string  `json:"nama"`
	NIK             string  `json:"nik"`
	Whatsapp        string  `json:"whatsapp"`
	Gender          string  `json:"gender"`
	Alamat          string  `json:"alamat"`
	Gaji            float64 `json:"gaji"`
	TanggalBergabung string `json:"tanggal_bergabung"`
}
