package controller

import (
	"fmt"
	"os"
)

// main function to delete book
func HapusBuku() {
	fmt.Println("=================================")
	fmt.Println("Hapus Buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	var kodeBuku string
	fmt.Print("Masukan Kode Buku : ")
	_, err := fmt.Scanln(&kodeBuku)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	// remove file json if kode buku exist
	err = os.Remove(fmt.Sprintf("books/book-%s.json", kodeBuku))
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		fmt.Printf("Buku dengan Kode: %s tidak ditemukan\n", kodeBuku)
	}

	fmt.Println("Buku Berhasil Dihapus!")

}
