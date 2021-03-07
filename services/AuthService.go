package services

import (
	"apitest/db"
	"apitest/models"
	"apitest/utils"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var login models.Login
	alert := utils.Alert{Msg: "Login สำเร็จ", Title: "Success", Type: "success"}
	if c.ShouldBind(&login) == nil {
		var queryUser models.User
		if login.Username != "" {
			log.Println("login with username")
			if err := db.GetDB().First(&queryUser, "username = ?", login.Username).Error; err != nil {
				alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
				utils.RespondJSON(c, 404, login, alert)
			} else {
				log.Println("user--->", queryUser)
				if checkPasswordHash(login.Password, queryUser.Password) == false {
					alert = utils.Alert{Msg: "invalid password", Title: "Error", Type: "error"}
					utils.RespondJSON(c, 404, login, alert)
				} else {
					jwttoken := models.Jwttoken{}
					jwttoken.Pubkey = "--public-key----"
					jwttoken.Reftoken = "---refresh token----"
					jwttoken.Token = "---token----"
					jwttoken.UserID = queryUser.ID
					jwttoken.Ulevel = queryUser.Level
					jwttoken.Roles = "Admin"
					utils.RespondJSON(c, 200, jwttoken, alert)
				}
			}
		} else if login.Email != "" {
			log.Println("login with email")
			if err := db.GetDB().First(&queryUser, "email = ?", login.Email).Error; err != nil {
				alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
				utils.RespondJSON(c, 404, login, alert)
			} else {
				if checkPasswordHash(login.Password, queryUser.Password) == false {
					alert = utils.Alert{Msg: "invalid password", Title: "Error", Type: "error"}
					utils.RespondJSON(c, 404, login, alert)
				} else {
					jwttoken := models.Jwttoken{}
					jwttoken.Pubkey = "--public-key----"
					jwttoken.Reftoken = "---refresh token----"
					jwttoken.Token = "---token----"
					jwttoken.UserID = queryUser.ID
					jwttoken.Ulevel = queryUser.Level
					jwttoken.Roles = "Admin"
					utils.RespondJSON(c, 200, jwttoken, alert)
				}
			}
		} else {
			log.Println("Error No Username/Email")
			alert = utils.Alert{Msg: "Error No Username/Email", Title: "Error", Type: "error"}
			utils.RespondJSON(c, 400, login, alert)
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func Register(c *gin.Context) {
	var user models.User
	alert := utils.Alert{Msg: "Register สำเร็จ", Title: "Success", Type: "success"}
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		if err := db.GetDB().Create(&user).Error; err != nil {
			alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
			utils.RespondJSON(c, 400, []string{}, alert)
		} else {
			user.Password = "---Hide-----"
			utils.RespondJSON(c, 400, user, alert)
		}
	} else {
		alert = utils.Alert{Msg: "unable to bind data", Title: "Error", Type: "error"}
		utils.RespondJSON(c, 400, []string{}, alert)
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
