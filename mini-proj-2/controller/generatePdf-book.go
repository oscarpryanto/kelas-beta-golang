package controller

import (
	"encoding/json"
	"fmt"
	"miniproject1/manajemen_buku_perpustakaan/model"
	"miniproject1/manajemen_buku_perpustakaan/utility"
	"os"

	"github.com/go-pdf/fpdf"
)

// main function to generate pdf
func GeneratePdfPesanan() {
	_ = os.Mkdir("pdf", 0777) //create directory
	var pilihanMenu int
	fmt.Println("=================================")
	fmt.Println("Generate Pdf file Buku")
	fmt.Println("=================================")

	// menu options to create pdf
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
	buku := model.Buku{}
	fmt.Println("=================================")
	fmt.Println("Generate Pdf sebuah buku")
	fmt.Println("=================================")
	TampilkanDaftarBuku()
	fmt.Println("=================================")

	fmt.Print("Masukan Kode Buku : ")
	kodeBuku := utility.InputMultiKata()

	dataJSON, err := os.ReadFile(fmt.Sprintf("books/book-%s.json", kodeBuku))
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		fmt.Printf("Buku dengan Kode: %s tidak ditemukan\n", kodeBuku)
		return
	}

	err = json.Unmarshal(dataJSON, &buku)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	textBook := fmt.Sprintf(
		"No.1 \n Kode buku: %s \n Judul Buku: %s \n Pengarang: %s \n Penerbit: %s \n Jumlah Halaman: %d \n Tahun Terbit: %d\n",
		buku.Kode,
		buku.Judul,
		buku.Pengarang,
		buku.Penerbit,
		buku.JumlahHalaman,
		buku.TahunTerbit)

	fileName := "buku-" + buku.Kode + ".pdf"
	fullPath := "pdf/" + fileName

	pdf.MultiCell(0, 10, textBook, "0", "L", false)
	pdf.Ln(5)
	err = pdf.OutputFileAndClose(fullPath)

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Printf("Berhasil Generate File: %s \n", fileName)
}

// function to create all data into pdf file
func generarteAllFile() {
	fmt.Println("=================================")
	fmt.Println("Generate all Book in Pdf....")
	fmt.Println("=================================")

	buku := LihatDaftarBuku()
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for i, item := range buku {
		textBook := fmt.Sprintf(
			"No.%d \n Kode buku: %s \n Judul Buku: %s \n Pengarang: %s \n Penerbit: %s \n Jumlah Halaman: %d \n Tahun Terbit: %d\n",
			i+1,
			item.Kode,
			item.Judul,
			item.Pengarang,
			item.Penerbit,
			item.JumlahHalaman,
			item.TahunTerbit)

		pdf.MultiCell(0, 10, textBook, "0", "L", false)
		pdf.Ln(5)
		fileName := "buku-" + item.Kode + ".pdf"
		fullPath := "pdf/" + fileName

		pdf.Ln(5)
		err := pdf.OutputFileAndClose(fullPath)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Printf("Berhasil Generate File: %s \n", fileName)
	}

}
