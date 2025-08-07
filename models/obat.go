package models

type Obat struct {
	ID           int    `json:"id"`
	NamaPenyakit string `json:"nama_penyakit"`
	NamaObat     string `json:"nama_obat"`
	Expired      string `json:"expired"`     // Format YYYY-MM-DD
	Keterangan   string `json:"keterangan"`
}