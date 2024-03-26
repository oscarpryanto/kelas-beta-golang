package utility

import (
	"miniproject3/manajemen_buku_perpustakaan/config"
	"miniproject3/manajemen_buku_perpustakaan/model"
	"time"
)

func GetBookList() ([]model.Book, error) {
	var books model.Book
	return books.GetAllData(config.Mysql.DB)

}

func GetBookById(id uint) (model.Book, error) {
	books := model.Book{
		Model: model.Model{
			ID: id,
		},
	}

	return books.GetById(config.Mysql.DB)
}

func GetBookByIdCount(id uint) (int64, error) {
	books := model.Book{
		Model: model.Model{
			ID: id,
		},
	}

	return books.GetByIdCount(config.Mysql.DB)
}

func InsertBookData(data model.Book) (uint, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	id, err := data.Create(config.Mysql.DB)
	return id, err
}

func DeleteBookById(id uint) error {
	books := model.Book{
		Model: model.Model{
			ID: id,
		},
	}

	return books.DeleteById(config.Mysql.DB)
}

func UpdateBookById(data model.Book) error {

	return data.UpdateOneById(config.Mysql.DB)
}
