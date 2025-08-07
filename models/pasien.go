package models

type Pasien struct {
	ID                int    `json:"id"`
	Nama              string `json:"nama_lengkap"`
	NIK               string `json:"nik"`
	Alamat            string `json:"alamat"`
	Telepon           string `json:"no_telepon"`
	BPJS              string `json:"bpjs"`
	Gender            string `json:"jenis_kelamin"`
	Keluhan           string `json:"keluhan"`
	TanggalRegistrasi string `json:"tanggal_registrasi"`
}