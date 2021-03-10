package services

import (
	"apitest/db"
	"apitest/models"
	"apitest/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var login models.Login
	payload := utils.GetResponse()
	if c.ShouldBind(&login) == nil {
		var queryUser models.User
		if login.Username != "" {
			log.Println("login with username")
			if err := db.GetDB().Preload("Roles").First(&queryUser, "username = ?", login.Username).Error; err != nil {
				alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
				payload.Alert = alert
				payload.Meta = http.StatusText(404)
				utils.RespondJSON(c, 404, payload)
			} else {
				log.Println("user--->", queryUser)
				if checkPasswordHash(login.Password, queryUser.Password) == false {
					alert := utils.Alert{Msg: "invalid password", Title: "Error", Type: "error"}
					payload.Alert = alert
					payload.Meta = http.StatusText(404)
					utils.RespondJSON(c, 404, payload)
				} else {
					jwttoken := models.Jwttoken{}
					// jwttoken.Token = utils.JwtSecretSign(queryUser)  // user secret key
					jwttoken.Token = utils.JwtSign(queryUser) // user private key
					jwttoken.Pubkey = utils.GetPubkey()
					jwttoken.Reftoken = queryUser.RememberToken
					jwttoken.UserID = queryUser.ID
					jwttoken.Ulevel = queryUser.Level
					jwttoken.Roles = queryUser.Roles
					payload.Data = jwttoken
					utils.RespondJSON(c, 200, payload)
				}
			}
		} else if login.Email != "" {
			log.Println("login with email")
			if err := db.GetDB().First(&queryUser, "email = ?", login.Email).Error; err != nil {
				alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
				payload.Alert = alert
				payload.Meta = http.StatusText(404)
				utils.RespondJSON(c, 404, payload)
			} else {
				if checkPasswordHash(login.Password, queryUser.Password) == false {
					alert := utils.Alert{Msg: "invalid password", Title: "Error", Type: "error"}
					payload.Alert = alert
					payload.Meta = http.StatusText(404)
					utils.RespondJSON(c, 404, payload)
				} else {
					jwttoken := models.Jwttoken{}
					// jwttoken.Token = utils.JwtSecretSign(queryUser)  // user secret key
					jwttoken.Token = utils.JwtSign(queryUser) // user private key
					jwttoken.Pubkey = utils.GetPubkey()
					jwttoken.Reftoken = queryUser.RememberToken
					jwttoken.UserID = queryUser.ID
					jwttoken.Ulevel = queryUser.Level
					jwttoken.Roles = queryUser.Roles
					payload.Data = jwttoken
					utils.RespondJSON(c, 200, payload)
				}
			}
		} else {
			log.Println("Error No Username/Email")
			alert := utils.Alert{Msg: "Error No Username/Email", Title: "Error", Type: "error"}
			payload.Alert = alert
			payload.Meta = http.StatusText(404)
			utils.RespondJSON(c, 404, payload)
		}
	} else {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	}
}

func Register(c *gin.Context) {
	var user models.User
	payload := utils.GetResponse()
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		if err := db.GetDB().Create(&user).Error; err != nil {
			alert := utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
			payload.Alert = alert
			payload.Meta = http.StatusText(404)
			utils.RespondJSON(c, 404, payload)
		} else {
			user.Password = "---Hide-----"
			payload.Data = user
			utils.RespondJSON(c, 200, payload)
		}
	} else {
		alert := utils.Alert{Msg: "unable to bind data", Title: "Error", Type: "error"}
		payload.Alert = alert
		payload.Meta = http.StatusText(404)
		utils.RespondJSON(c, 404, payload)
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
