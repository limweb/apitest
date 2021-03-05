package services

import (
	"apitest/db"
	"apitest/models"
	"fmt"
)

func GetAllBook(b *[]models.Book) (err error) {
	if err = db.GetDB().Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(b *models.Book) (err error) {
	if err = db.GetDB().Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *models.Book, id string) (err error) {
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *models.Book, id string) (err error) {
	fmt.Println(b)
	db.GetDB().Save(b)
	return nil
}

func DeleteBook(b *models.Book, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(b)
	return nil
}
