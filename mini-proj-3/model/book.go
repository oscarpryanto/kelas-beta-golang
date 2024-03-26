package model

import (
	"gorm.io/gorm"
)

type Book struct {
	Model
	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

func (cr *Book) Create(db *gorm.DB) (uint, error) {
	err := db.Model(Book{}).Create(&cr).Error

	if err != nil {
		return cr.ID, err
	}

	return cr.ID, nil
}

func (cr *Book) GetById(db *gorm.DB) (Book, error) {
	res := Book{}
	err := db.Model(Book{}).Where("id = ?", cr.Model.ID).Take(&res).Error

	if err != nil {
		return Book{}, err
	}

	return res, nil
}

func (cr *Book) GetByIdCount(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(Book{}).Where("id = ?", cr.Model.ID).Count(&count).Error

	if err != nil {
		return count, err
	}

	return count, nil
}

func (cr *Book) GetAllData(db *gorm.DB) ([]Book, error) {
	res := []Book{}
	err := db.Model(Book{}).Find(&res).Error

	if err != nil {
		return []Book{}, err
	}

	return res, nil
}

func (cr *Book) UpdateOneById(db *gorm.DB) error {
	err := db.Model(Book{}).Select("isbn",
		"penulis",
		"tahun",
		"judul",
		"gambar",
		"stok").Where("id = ?", cr.Model.ID).Updates(map[string]interface{}{
		"isbn":    cr.ISBN,
		"penulis": cr.Penulis,
		"tahun":   cr.Tahun,
		"judul":   cr.Judul,
		"gambar":  cr.Gambar,
		"stok":    cr.Stok,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Book) DeleteById(db *gorm.DB) error {
	err := db.Model(Book{}).Where("id = ?", cr.Model.ID).Delete(&cr).Error

	if err != nil {
		return err
	}

	return nil
}
