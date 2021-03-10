package db

import (
	"apitest/config"
	"apitest/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	dB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

func SetupDB() {
	db := dB
	config := config.GetConfig()
	driver := config.Database.Driver
	database := config.Database.Dbname
	username := config.Database.Username
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port

	if driver == "sqlite" {
		db, err = gorm.Open("sqlite3", "./webapi.db")
		if err != nil {
			DBErr = err
			fmt.Println("db err: ", err)
		}
	} else if driver == "postgres" {
		db, err = gorm.Open("postgres", "host="+host+" port="+port+" user="+username+" dbname="+database+"  sslmode=disable password="+password)
		if err != nil {
			DBErr = err
			fmt.Println("db err: ", err)
		}
	} else if driver == "mysql" {
		db, err = gorm.
			Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			DBErr = err
			fmt.Println("db err: ", err)
		}
	}

	// Change this to true if you want to see SQL queries
	// ไว้แสดง Sql ออกทาง Log เมื่อต้องการ Debug คำสั่ง Sql
	db.LogMode(config.Database.LogMode)

	//----------------Add MigrateDB --------------------------------
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.PermissionRole{})
	db.AutoMigrate(&models.RoleUser{})
	db.AutoMigrate(&models.PasswordResets{})

	dB = db
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return dB
}

// GetDBErr helps you to get a connection
func GetDBErr() error {
	return DBErr
}
