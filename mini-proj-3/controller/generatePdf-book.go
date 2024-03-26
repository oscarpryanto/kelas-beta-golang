package controller

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/utility"
	"os"

	"github.com/go-pdf/fpdf"
)

func GeneratePdfBuku() {
	_ = os.Mkdir("pdf", 0777) //create directory
	var pilihanMenu int
	fmt.Println("=================================")
	fmt.Println("Generate Pdf file Buku")
	fmt.Println("=================================")

	for {
		fmt.Println("=================================")
		fmt.Println("Silahkan Pilih menu pada fitur Generate Pdf : ")
		fmt.Println("1. Pilih buku yang ingin di Generate")
		fmt.Println("2. Generate seluruh buku")
		fmt.Println("3. Keluar/batal")
		fmt.Println("===========================================")
		fmt.Print("Masukan Pilihan : ")
		_, err := fmt.Scanln(&pilihanMenu)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		switch pilihanMenu {
		case 1:
			generarteSingleFile()
		case 2:
			generarteAllFile()
		case 3:
			return
		}

	}
}

// function to create single pdf
func generarteSingleFile() {

	fmt.Println("=================================")
	fmt.Println("Generate Pdf sebuah buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	fmt.Print("Masukan ID Buku : ")
	var inputanId uint
	_, err := fmt.Scanln(&inputanId)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	dataBook, err := utility.GetBookById(inputanId)
	if err != nil {
		fmt.Println("data buku tidak ditemukan", err)
		return
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	textBook := fmt.Sprintf("No.%d | ID: %d | ISBN: %s | Judul Buku: %s | Penulis: %s | Tahun Terbit: %d | Stok: %d | Gambar: %s\n",
		1,
		dataBook.ID,
		dataBook.ISBN,
		dataBook.Judul,
		dataBook.Penulis,
		dataBook.Tahun,
		dataBook.Stok,
		dataBook.Gambar,
	)

	fileName := "buku-" + dataBook.ISBN + ".pdf"
	fullPath := "pdf/" + fileName

	pdf.MultiCell(0, 10, textBook, "0", "L", false)
	pdf.Ln(5)
	err = pdf.OutputFileAndClose(fullPath)

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Printf("Berhasil Generate File: %s \n", fileName)
}

func generarteAllFile() {
	fmt.Println("=================================")
	fmt.Println("Generate all Book in Pdf....")
	fmt.Println("=================================")

	// buku := LihatDaftarBuku()
	listBook, err := utility.GetBookList()
	if err != nil {
		fmt.Println("Terjadi kesalahan pada saat mengambil data buku")
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for urutan, item := range listBook {
		textBook := fmt.Sprintf("No.%d | ID: %d | ISBN: %s | Judul Buku: %s | Penulis: %s | Tahun Terbit: %d | Stok: %d | Gambar: %s\n",
			urutan+1,
			item.ID,
			item.ISBN,
			item.Judul,
			item.Penulis,
			item.Tahun,
			item.Stok,
			item.Gambar,
		)

		fileName := "buku-" + item.ISBN + ".pdf"
		fullPath := "pdf/" + fileName
		pdf.MultiCell(0, 10, textBook, "0", "L", false)
		pdf.Ln(5)

		err := pdf.OutputFileAndClose(fullPath)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Printf("Berhasil Generate File: %s \n", fileName)
	}

}
