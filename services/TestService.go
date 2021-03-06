package services

import (
	"apitest/db"
	"apitest/models"
	"fmt"
) 
 
func GetAllTest(b *[]models.Test) (err error) { 
	if err = db.GetDB().Find(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func GetAllIdDescTest(b *[]models.Test) (err error) { 
	if err = db.GetDB().Order("id desc").Find(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func AddNewTest(b *models.Test) (err error) { 
	if err = db.GetDB().Create(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func GetOneTest(b *models.Test, id string) (err error) { 
	if err := db.GetDB().Where("id = ?", id).First(b).Error; err != nil { 
		return err 
	} 
	return nil 
} 
 
func PutOneTest(b *models.Test, id string) (err error) { 
	fmt.Println(b) 
	db.GetDB().Save(b) 
	return nil 
} 
 
func DeleteTest(b *models.Test, id string) (err error) { 
	db.GetDB().Where("id = ?", id).Delete(b) 
	return nil 
} 
 