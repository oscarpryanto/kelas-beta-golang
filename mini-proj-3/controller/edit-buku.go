package controller

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/utility"

	"github.com/sirupsen/logrus"
)

func EditBuku() {
	// buku := model.Book{}
	var pilihanMenu int
	var inputanTahun uint
	var inputanStok uint

	listBook, _ := utility.GetBookList()
	if len(listBook) <= 0 {
		fmt.Println("Data buku masih kosong!")
		fmt.Println("Silahkan tambahkan data buku terlebih dahulu!")
		return
	}

	fmt.Println("=================================")
	fmt.Println("Ubah Data Buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	fmt.Print("Masukan ID buku yang ingin dihapus : ")
	var inputanId uint
	_, err := fmt.Scanln(&inputanId)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	dataBook, err := utility.GetBookById(inputanId)
	if err != nil {

		logrus.Error("Error on get car data", err.Error())
		fmt.Println("data buku tidak ditemukan", err)

	}

	utility.ClearTerminal()
	for {
		fmt.Println("=================================")
		fmt.Println("Silahkan Pilih elemen data pada buku yang ingin diubah : ")
		fmt.Println("1. ISBN")
		fmt.Println("2. Judul")
		fmt.Println("3. Penulis")
		fmt.Println("4. Tahun")
		fmt.Println("5. Gambar")
		fmt.Println("6. Stok")
		fmt.Println("7. Simpan Perubahan")
		fmt.Println("8. Keluar/batal")
		fmt.Println("===========================================")
		fmt.Print("Masukan Pilihan : ")

		_, err = fmt.Scanln(&pilihanMenu)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		switch pilihanMenu {
		case 1:
			fmt.Print("Silahkan Masukan ISBN Buku : ")
			dataBook.ISBN = utility.InputMultiKata()
			fmt.Println("ISBN Buku berhasil diubah ")
		case 2:
			fmt.Print("Silahkan Masukan Judul Buku : ")
			dataBook.Judul = utility.InputMultiKata()
			fmt.Println("Judul Buku berhasil diubah ")
		case 3:
			fmt.Print("Silahkan Masukan Penulis Buku : ")
			dataBook.Penulis = utility.InputMultiKata()
			fmt.Println("Penulis Buku berhasil diubah ")
		case 4:
			fmt.Print("Silahkan Masukan Tahun Terbit Buku : ")
			_, error := fmt.Scanln(&inputanTahun)
			if error != nil {
				fmt.Print("Terjadi kesalahan pada saat edit tahun buku : ", error)
				return
			}
			dataBook.Tahun = inputanTahun
			fmt.Println("Tahun terbit buku berhasil diubah ")
		case 5:
			fmt.Print("Silahkan Masukan Gambar Buku : ")
			dataBook.Gambar = utility.InputMultiKata()
			fmt.Println("Gambar Buku berhasil diubah ")
		case 6:
			fmt.Print("Silahkan Masukan Stok Buku : ")
			_, error := fmt.Scanln(&inputanStok)
			if error != nil {
				fmt.Print("Terjadi kesalahan pada saat edit stok buku : ", error)
				return
			}
			dataBook.Stok = inputanStok
			fmt.Println("Stok buku berhasil diubah ")
		}

		utility.ClearTerminal()
		if pilihanMenu == 7 {
			break
		}
		if pilihanMenu == 8 {
			return
		}

	}

	fmt.Println("Mengupdate data buku .....")
	err = utility.UpdateBookById(dataBook)
	if err != nil {
		fmt.Println("Terjadi error pada saat mengupdate data:", err)
	}

	fmt.Printf("Berhasil menyimpan perubahan data buku dengan ID: %d\n", dataBook.ID)
	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")

}
