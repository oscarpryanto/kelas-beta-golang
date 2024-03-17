package controller

import (
	"encoding/json"
	"fmt"
	"miniproject1/manajemen_buku_perpustakaan/model"
	"miniproject1/manajemen_buku_perpustakaan/utility"
	"os"
)

// main function untuk ubah buku
func UbahBuku() {
	var pilihanMenu int
	inputJumlahHalaman := 0
	inputTahunTerbit := 0
	buku := model.Buku{}

	fmt.Println("=================================")
	fmt.Println("Ubah Data Buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	fmt.Print("Masukan kode Buku yang akan diubah : ")
	inputanUbahKodeBuku := utility.InputMultiKata()

	// read file berdasarkan kode yang sudah di input
	dataJSON, err := os.ReadFile(fmt.Sprintf("books/book-%s.json", inputanUbahKodeBuku))
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		fmt.Printf("Buku dengan Kode: %s tidak ditemukan\n", inputanUbahKodeBuku)
		return
	}

	// parse data json into struct buku
	err = json.Unmarshal(dataJSON, &buku)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	// utility to clear terminal
	utility.ClearTerminal()

	// menu for editing every elemnet of book
	for {
		fmt.Println("=================================")
		fmt.Println("Silahkan Pilih elemen data pada buku yang ingin diubah : ")
		fmt.Println("1. Judul Buku")
		fmt.Println("2. Pengarang")
		fmt.Println("3. PenerBit")
		fmt.Println("4. Jumlah Halaman")
		fmt.Println("5. Tahun Terbit")
		fmt.Println("6. Simpan Perubahan")
		fmt.Println("7. Keluar/batal")
		fmt.Println("===========================================")
		fmt.Print("Masukan Pilihan : ")

		_, err = fmt.Scanln(&pilihanMenu)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		switch pilihanMenu {
		case 1:
			fmt.Print("Silahkan Masukan Judul Buku : ")
			inputanJudulBuku := utility.InputMultiKata()
			buku.Judul = inputanJudulBuku
			fmt.Println("Judul Buku berhasil diubah ")
		case 2:
			fmt.Print("Silahkan Masukan Pengarang Buku : ")
			inputanPengarangBuku := utility.InputMultiKata()
			buku.Pengarang = inputanPengarangBuku
			fmt.Println("Pengarang Buku berhasil diubah ")
		case 3:
			fmt.Print("Silahkan Masukan Penerbit Buku : ")
			inputanPenerbitBuku := utility.InputMultiKata()
			buku.Penerbit = inputanPenerbitBuku
			fmt.Println("Penerbit Buku berhasil diubah ")
		case 4:
			fmt.Print("Silahkan Masukan Jumlah Halaman Buku : ")
			_, error := fmt.Scanln(&inputJumlahHalaman)
			if error != nil {
				fmt.Print("Terjadi kesalahan pada saat menambahkan jumlah halaman buku : ", error)
				return
			}
			buku.JumlahHalaman = inputJumlahHalaman
			fmt.Println("Jumlah halam buku berhasil diubah ")
		case 5:
			fmt.Print("Silahkan Masukan Tahun terbit Buku : ")
			_, error := fmt.Scanln(&inputTahunTerbit)
			if error != nil {
				fmt.Print("Terjadi kesalahan pada saat menambahkan Tahun terbit buku : ", error)
				return
			}
			buku.TahunTerbit = inputTahunTerbit
			fmt.Println("Jumlah tahun terbit buku berhasil diubah ")
		}

		utility.ClearTerminal()

		if pilihanMenu == 6 {
			break
		}
		if pilihanMenu == 7 {
			return
		}

	}

	// parse data struct into json format
	dataJson2, err := json.MarshalIndent(buku, "", "	")
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	// write file into folder and add file name formating
	err = os.WriteFile(fmt.Sprintf("books/book-%s.json", buku.Kode), dataJson2, 0644)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Printf("Berhasil menyimpan perubahan data buku dengan kode: %s\n", buku.Kode)
	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	// display data after editing
	fmt.Printf("No.1 | Kode buku: %s | Judul Buku: %s | Pengarang: %s | Penerbit: %s | Jumlah Halaman: %d | Tahun Terbit: %d\n",
		buku.Kode,
		buku.Judul,
		buku.Pengarang,
		buku.Penerbit,
		buku.JumlahHalaman,
		buku.TahunTerbit,
	)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------------------------------------------")

}
