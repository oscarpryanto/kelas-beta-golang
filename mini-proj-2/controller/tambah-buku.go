package controller

import (
	"encoding/json"
	"fmt"
	"miniproject1/manajemen_buku_perpustakaan/model"
	"miniproject1/manajemen_buku_perpustakaan/utility"
	"os"
)

// main function of Tambah data buku
func TambahBuku() {

	inputJumlahHalaman := 0
	inputTahunTerbit := 0

	fmt.Println("===========================================")
	fmt.Println("Tambah Data Buku")
	fmt.Println("===========================================")

	draftBuku := []model.Buku{}

	existingBook := LihatDaftarBuku()

	// loop until finish add data buku
	for {
		// input kode
		var inputKodeBuku string
		
		// looping to check every kode buku is already exist in draft and json file
		for {
			breakOp := 1
			fmt.Print("Silahkan Masukan Kode Buku : ")
			inputKodeBuku = utility.InputMultiKata()

			if len(draftBuku) > 0 {
				for _, cekIfKodeBukuDraftAlreadyExist := range draftBuku {
					if cekIfKodeBukuDraftAlreadyExist.Kode == inputKodeBuku {
						fmt.Println("Kode buku pada draft sudah ada, silahkan masukan kode lain")
						breakOp = 0

						break
					}
				}
			}

			for _, cekIfKodeBookAlreadyExist := range existingBook {
				if cekIfKodeBookAlreadyExist.Kode == inputKodeBuku {
					fmt.Println("Kode buku sudah ada, silahkan masukan kode lain")
					breakOp = 0
					break
				}
			}

			if breakOp == 1 {
				break
			}
		}

		// input judul
		fmt.Print("Silahkan Masukan Judul Buku : ")
		inputanJudulBuku := utility.InputMultiKata()

		// input pengarang
		fmt.Print("Silahkan Masukan Pengarang Buku : ")
		inputanPengarangBuku := utility.InputMultiKata()

		// input penerbit
		fmt.Print("Silahkan Masukan Penerbit Buku : ")
		inputanPenerbitBuku := utility.InputMultiKata()

		// input jumlah halaman
		fmt.Print("Silahkan Masukan Jumlah Halaman Buku : ")
		_, error := fmt.Scanln(&inputJumlahHalaman)
		if error != nil {
			fmt.Print("Terjadi kesalahan pada saat menambahkan jumlah halaman buku : ", error)
			return
		}

		// input tahun terbit
		fmt.Print("Silahkan Masukan Tahun Terbit Buku : ")
		_, error = fmt.Scanln(&inputTahunTerbit)
		if error != nil {
			fmt.Print("Terjadi kesalahan pada saat menambahkan tahun terbit buku : ", error)
			return
		}

		draftBuku = append(draftBuku, model.Buku{
			Kode:          inputKodeBuku,
			Judul:         inputanJudulBuku,
			Pengarang:     inputanPengarangBuku,
			Penerbit:      inputanPenerbitBuku,
			JumlahHalaman: inputJumlahHalaman,
			TahunTerbit:   inputTahunTerbit,
		})

		var pilihanMenuBuku = 0
		fmt.Println("Ketik 1 untuk tambah data buku, ketik 0 untuk keluar")
		_, error = fmt.Scanln(&pilihanMenuBuku)
		if error != nil {
			fmt.Println("Terjadi Error:", error)
			return
		}

		if pilihanMenuBuku == 0 {
			break
		}

	}

	fmt.Println("Menambah data buku .....")
	_ = os.Mkdir("books", 0777) //create directory

	for _, itemBuku := range draftBuku {
		dataJson, err := json.MarshalIndent(itemBuku, "", "	")
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		err = os.WriteFile(fmt.Sprintf("books/book-%s.json", itemBuku.Kode), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Printf("Berhasil menyimpan data buku dengan kode: %s\n", itemBuku.Kode)
	}

}
