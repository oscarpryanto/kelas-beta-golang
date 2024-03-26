package controller

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/model"
	"miniproject3/manajemen_buku_perpustakaan/utility"

	"github.com/sirupsen/logrus"
)

func TambahBuku() {

	fmt.Println("===========================================")
	fmt.Println("Tambah Data Buku")
	fmt.Println("===========================================")

	var inputanTahun uint
	var inputanStok uint
	draftBuku := []model.Book{}

	for {

		// input ISBN
		fmt.Print("Silahkan Masukan ISBN Buku : ")
		inputanIsbn := utility.InputMultiKata()

		// input Penulis
		fmt.Print("Silahkan Masukan Penulis Buku : ")
		inputanPenulis := utility.InputMultiKata()

		// input judul
		fmt.Print("Silahkan Masukan Judul Buku : ")
		inputanJudul := utility.InputMultiKata()

		// input tahun
		fmt.Print("Silahkan Masukan Tahun Terbit Buku : ")
		_, error := fmt.Scanln(&inputanTahun)
		if error != nil {
			fmt.Print("Terjadi kesalahan pada saat menambahkan tahun terbit buku : ", error)
			return
		}

		// input gambar
		fmt.Print("Silahkan Masukan Gambar Buku : ")
		inputanGambar := utility.InputMultiKata()

		// input stok
		fmt.Print("Silahkan Masukan Jumlah Stok Buku : ")
		_, error = fmt.Scanln(&inputanStok)
		if error != nil {
			fmt.Print("Terjadi kesalahan pada saat menambahkan stok buku : ", error)
			return
		}

		draftBuku = append(draftBuku, model.Book{
			ISBN:    inputanIsbn,
			Penulis: inputanPenulis,
			Tahun:   inputanTahun,
			Judul:   inputanJudul,
			Gambar:  inputanGambar,
			Stok:    inputanStok,
		})

		var pilihanMenuBuku = 0
		fmt.Println("Ketik 1 untuk tambah data buku, ketik 0 untuk keluar")
		_, err := fmt.Scanln(&pilihanMenuBuku)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		if pilihanMenuBuku == 0 {
			break
		}
	}

	fmt.Println("Menambah data buku .....")
	for _, itemBuku := range draftBuku {
		id, err := utility.InsertBookData(itemBuku)
		if err != nil {
			logrus.Error("Error when post data", err.Error())
			fmt.Println("Dengan ID ISBN = ", itemBuku.ISBN)
		}
		fmt.Println("Sukses menambahkan buku dengan id =", id)
	}

}
