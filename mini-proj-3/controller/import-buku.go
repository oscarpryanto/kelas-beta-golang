package controller

import (
	"encoding/csv"
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/model"
	"miniproject3/manajemen_buku_perpustakaan/utility"
	"os"
	"strconv"
)

func ImportDataBuku() {
	fmt.Println("=================================")
	fmt.Println("Import Csv data Buku")
	fmt.Println("=================================")

	fmt.Println("")
	fmt.Print("Masukan Pilihan lokasi file csv : ")
	inputanLokasiFile := utility.InputMultiKata()

	fileOpen, err := os.Open(inputanLokasiFile)
	if err != nil {
		fmt.Println("terjadi error pada saat membukan file = ", err)
		return
	}

	readerFile := csv.NewReader(fileOpen)
	recordsFile, err := readerFile.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	books := csvToStruct(recordsFile)
	totalRow := 0
	totalInsert := 0
	for index, book := range books {
		_, err := utility.GetBookByIdCount(book.ID)
		indexInsert := 0

		if err != nil {
			indexInsert++
			totalInsert = indexInsert
			_, err := utility.InsertBookData(book)
			if err != nil {
				fmt.Println("error menambahkan buku ID = ", book.ID)
			}
		}

		err = utility.UpdateBookById(book)
		if err != nil {
			fmt.Println("Terjadi error pada saat mengupdate data:", book.ID)
		}

		index++
		totalRow = index
	}

	utility.ClearTerminal()
	fmt.Println("Selesai melakukan import data buku!")
	fmt.Printf("Berhasil mengimport %d row data\n", totalRow)
	fmt.Printf(" %d Berhasil di insert\n", totalInsert)
	fmt.Printf(" %d Berhasil di update\n", totalRow-totalInsert)
	defer fileOpen.Close()

}

func csvToStruct(records [][]string) []model.Book {
	books := []model.Book{}

	for _, book := range records {
		idStr := book[0]
		tahunStr := book[3]
		stokStr := book[6]

		id, errId := strconv.ParseUint(idStr, 10, 32)
		if errId != nil {
			fmt.Println("Error converting ID:", errId)
			continue
		}
		tahun, errTahun := strconv.ParseUint(tahunStr, 10, 32)
		if errTahun != nil {
			fmt.Println("Error converting tahun:", errTahun)
			continue
		}

		stok, errStok := strconv.ParseUint(stokStr, 10, 32)
		if errStok != nil {
			fmt.Println("Error converting stok:", errStok)
			continue
		}
		books = append(books, model.Book{
			Model: model.Model{
				ID: uint(id),
			},
			ISBN:    book[1],
			Penulis: book[2],
			Tahun:   uint(tahun),
			Judul:   book[4],
			Gambar:  book[5],
			Stok:    uint(stok),
		})
	}

	return books
}
