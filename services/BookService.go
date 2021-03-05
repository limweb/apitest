package services

import (
	"apitest/db"
	"apitest/model"
	"fmt"
)

func GetAllBook(b *[]model.Book) (err error) {
	if err = db.GetDB().Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(b *model.Book) (err error) {
	if err = db.GetDB().Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *model.Book, id string) (err error) {
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *model.Book, id string) (err error) {
	fmt.Println(b)
	db.GetDB().Save(b)
	return nil
}

func DeleteBook(b *model.Book, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(b)
	return nil
}
