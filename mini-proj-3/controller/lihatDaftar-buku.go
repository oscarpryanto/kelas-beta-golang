package controller

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/utility"

	"github.com/sirupsen/logrus"
)

func TampilkanDaftarBuku() {
	fmt.Println("=================================")
	fmt.Println("Lihat Daftar Buku")
	fmt.Println("=================================")
	fmt.Println("Memuat data ...")

	listBook, err := utility.GetBookList()
	if err != nil {
		logrus.Error("Error on get cars list", err.Error())
		fmt.Println("Terjadi kesalahan pada saat mengambil data buku")
	}

	if len(listBook) <= 0 {
		fmt.Println("Data buku masih kosong!")
		fmt.Println("Silahkan tambahkan data buku terlebih dahulu!")
		return
	}

	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	for urutan, item := range listBook {
		fmt.Printf("No.%d | ID: %d | ISBN: %s | Judul Buku: %s | Penulis: %s | Tahun Terbit: %d | Stok: %d | Gambar: %s\n",
			urutan+1,
			item.ID,
			item.ISBN,
			item.Judul,
			item.Penulis,
			item.Tahun,
			item.Stok,
			item.Gambar,
		)
		fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	}
}
