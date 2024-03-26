package main

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/config"
	"miniproject3/manajemen_buku_perpustakaan/controller"
	"miniproject3/manajemen_buku_perpustakaan/utility"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	InitEnv()
	config.OpenDb()

	var pilihanMenu int
	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah data Buku")
	fmt.Println("2. Lihat daftar buku")
	fmt.Println("3. Ubah data buku")
	fmt.Println("4. Hapus data buku")
	fmt.Println("5. Import data buku")
	fmt.Println("6. Generate Pdf")
	fmt.Println("7. Keluar")
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
		controller.EditBuku()
	case 4:
		controller.Hapusbuku()
	case 5:
		controller.ImportDataBuku()
	case 6:
		controller.GeneratePdfBuku()
	case 7:
		os.Exit(0)
	}

	main()
}
