package controller

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/utility"

	"github.com/sirupsen/logrus"
)

func Hapusbuku() {
	listBook, _ := utility.GetBookList()
	if len(listBook) <= 0 {
		fmt.Println("Data buku masih kosong!")
		fmt.Println("Silahkan tambahkan data buku terlebih dahulu!")
		return
	}

	fmt.Println("=================================")
	fmt.Println("Hapus Buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	fmt.Print("Masukan ID buku yang ingin dihapus : ")
	var inputanId uint
	_, err := fmt.Scanln(&inputanId)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	err = utility.DeleteBookById(inputanId)
	if err != nil {
		logrus.Error("Error when post data", err.Error())
		fmt.Println("Dengan ID = ", inputanId)
	}
	fmt.Println("Sukses menghapus data buku dengan id = ", inputanId)
}
