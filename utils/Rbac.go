package utils

import (
	"apitest/db"
	"log"
)

type NameResult struct {
	Name string
}

type Permission struct {
	Name string `json:"name"  example:"Aaa" `                 // Rbac name
	Slug string `json:"slug" gorm:"unique" example:"Create" ` // Rbac slug
}

//get Role by UserID
func GetRoles(id uint) []string {
	var nameResults []NameResult
	db.GetDB().Raw("select  name from roles where id in ( SELECT role_id as id  FROM role_users  WHERE user_id = ? )", 1).Scan(&nameResults)
	log.Println("roles-->", nameResults)

	var result []string
	for _, v := range nameResults {
		result = append(result, v.Name)
	}
	return result
}

//get permission name by user id
func GetPermissions(id uint) []string {
	var nameResults []NameResult
	db.GetDB().Raw("select DISTINCT name from permissions where id in (	select permission_id as id from permission_role where role_id in (			select role_id as id from role_users where user_id = ?	))", 1).Scan(&nameResults)
	var result []string
	for _, v := range nameResults {
		result = append(result, v.Name)
	}
	return result
}

// get permission name and slug by user id
func GetPermislug(id uint) []Permission {
	var nameResults []Permission
	db.GetDB().Raw(`select DISTINCT name,slug from permissions where id in (
			select permission_id as id from permission_role where role_id in (
			select role_id as id from role_users where user_id = ?))
			`, 1).Scan(&nameResults)
	return nameResults
}
