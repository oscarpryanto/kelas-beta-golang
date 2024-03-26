package model_test

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/config"
	"miniproject3/manajemen_buku_perpustakaan/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func InitEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	config.OpenDb()
}

var dataBuku model.Book

func TestCreateBook(t *testing.T) {
	InitEnv()
	dataBuku.ISBN = "1233123"
	dataBuku.Judul = "Buku sekolah beta"
	dataBuku.Penulis = "Markas bali"
	dataBuku.Gambar = "https://blog.tripcetera.com/id/wp-content/uploads/2020/10/Danau-Toba-edited.jpg"
	dataBuku.Stok = 1
	dataBuku.Tahun = 2024

	id, err := dataBuku.Create(config.Mysql.DB)
	assert.Nil(t, err, "cek error create book")
	assert.NotEmpty(t, id, "cek id is not empty")

}

func TestGetBookById(t *testing.T) {
	InitEnv()
	res, err := dataBuku.GetById(config.Mysql.DB)
	assert.Nil(t, err, "cek error get data buku by id")
	assert.NotEmpty(t, res, "data buku kosong")

}

func TestGetAllBook(t *testing.T) {
	InitEnv()

	res, err := dataBuku.GetAllData(config.Mysql.DB)
	assert.Nil(t, err, "error saat mengambil data")
	assert.NotEmpty(t, res, "data buku kosong")
}

func TestUpdateBook(t *testing.T) {
	InitEnv()

	dataBuku.ISBN = "1234567"
	dataBuku.Judul = "Buku sekolah beta update"
	dataBuku.Penulis = "Markas bali"
	dataBuku.Gambar = "https://blog.tripcetera.com/id/wp-content/uploads/2020/10/Danau-Toba-edited.jpg"
	dataBuku.Stok = 2
	dataBuku.Tahun = 2023

	err := dataBuku.UpdateOneById(config.Mysql.DB)
	assert.Nil(t, err, "error update buku")

	assert.NotEqual(t, dataBuku.ISBN, "1233123")
	assert.NotEqual(t, dataBuku.Judul, "Buku sekolah beta")
	assert.Equal(t, dataBuku.Penulis, "Markas bali")
	assert.Equal(t, dataBuku.Gambar, "https://blog.tripcetera.com/id/wp-content/uploads/2020/10/Danau-Toba-edited.jpg")
	assert.NotEqual(t, dataBuku.Stok, 2)
	assert.NotEqual(t, dataBuku.Tahun, 2024)

}

func TestDeleteByID(t *testing.T) {
	InitEnv()
	err := dataBuku.DeleteById(config.Mysql.DB)
	assert.Nil(t, err, "error delete buku")

	res, err := dataBuku.GetByIdCount(config.Mysql.DB)
	assert.Nil(t, err, "cek error get data buku by id")
	assert.Empty(t, res)

}
