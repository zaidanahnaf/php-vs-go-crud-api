package models

type UMKM struct {
	ID        int    `json:"id"`
	NamaUsaha string `json:"nama_usaha"`
	Pemilik   string `json:"pemilik"`
	Kategori  string `json:"kategori"`
}
