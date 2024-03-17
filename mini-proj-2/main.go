package main

import (
	"fmt"
	"miniproject1/manajemen_buku_perpustakaan/controller"
	"miniproject1/manajemen_buku_perpustakaan/utility"
	"os"
)

func main() {
	var pilihanMenu int
	_ = os.Mkdir("books", 0777) //create directory
	_ = os.Mkdir("pdf", 0777)   //create directory

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah data Buku")
	fmt.Println("2. Lihat daftar buku")
	fmt.Println("3. Ubah data buku")
	fmt.Println("4. Hapus data buku")
	fmt.Println("5. Generate Pdf")
	fmt.Println("6. Keluar")
	fmt.Println("===========================================")
	fmt.Print("Masukan Pilihan : ")

	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
	utility.ClearTerminal()

	switch pilihanMenu {
	case 1:
		controller.TambahBuku()
	case 2:
		controller.TampilkanDaftarBuku()
	case 3:
		controller.UbahBuku()
	case 4:
		controller.HapusBuku()
	case 5:
		controller.GeneratePdfPesanan()
	case 6:
		os.Exit(0)
	}

	main()
}
