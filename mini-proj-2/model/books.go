package model

type Buku struct {
	Kode          string `json:"kode"`
	Judul         string `json:"judul"`
	Pengarang     string `json:"pengarang"`
	Penerbit      string `json:"penerbit"`
	JumlahHalaman int    `json:"jumlahHalaman"`
	TahunTerbit   int    `json:"tahunTerbit"`
}
