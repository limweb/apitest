package services

import (
	"apitest/db"
	"apitest/models"
	"fmt"
)

func GetAllOrder(b *[]models.Order) (err error) {
	if err = db.GetDB().Find(b).Error; err != nil {
		return err
	}
	return nil
}

func GetAllIdDescOrder(b *[]models.Order) (err error) {
	if err = db.GetDB().Order("id desc").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewOrder(b *models.Order) (err error) {
	if err = db.GetDB().Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneOrder(b *models.Order, id string) (err error) {
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneOrder(b *models.Order, id string) (err error) {
	fmt.Println(b)
	db.GetDB().Save(b)
	return nil
}

func DeleteOrder(b *models.Order, id string) (err error) {
	db.GetDB().Where("id = ?", id).Delete(b)
	return nil
}
