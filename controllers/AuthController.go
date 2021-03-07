package controllers

// controllers.SetupAuthRoutes(v1)
import (
	"apitest/services"

	"github.com/gin-gonic/gin"
)

func SetupLoginRoutes(rg *gin.Engine) {
	rg.POST("login", Login)
	rg.POST("register", Register)
}

func SetupAuthRoutes(rg *gin.Engine) {
	// rg.Use(utils.JwtSecretVerify)
	// rg.Use(utils.JwtPubkVerify)
	rg.GET("refreshtoken", Refreshtoken)
	rg.POST("resetpass", Resetpass)
	rg.GET("emailverify", Emailverify)
}

// Login godoc
// @summary Login Auths
// @description Login with username/email  and password
// @tags auths
// @id Login
// @accept json
// @produce json
// @Param login body models.Login true "Login info"
// @response 200 {object} []models.User "OK"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @Router /login [post]
func Login(c *gin.Context) {
	services.Login(c)
}

// Register godoc
// @summary Register Auths
// @description Register User
// @tags auths
// @id Register
// @accept json
// @produce json
// @Param login body models.User true "User info"
// @response 200 {array} []models.Login "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @Router /register [post]
func Register(c *gin.Context) {
	services.Register(c)
}

// Refreshtoken godoc
// @summary List Auths
// @description List all auths
// @tags auths
// @security ApiKeyAuth
// @id Refreshtoken
// @accept json
// @produce json
// @response 200 {array} []models.Login "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /refreshtoken [get]
func Refreshtoken(c *gin.Context) {
	// var rs []models.Login
	// err := services.GetAllAuth(&rs)
	// alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	// if err != nil {
	// 	alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
	// 	utils.RespondJSON(c, 404, []string{}, alert)
	// } else {
	// 	utils.RespondJSON(c, 200, rs, alert)
	// }
}

// Resetpass godoc
// @summary List Auths
// @description List all auths
// @tags auths
// @security ApiKeyAuth
// @id Resetpass
// @accept json
// @produce json
// @response 200 {array} []models.Login "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /resetpass [post]
func Resetpass(c *gin.Context) {
	// var rs []models.Login
	// err := services.GetAllAuth(&rs)
	// alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	// if err != nil {
	// 	alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
	// 	utils.RespondJSON(c, 404, []string{}, alert)
	// } else {
	// 	utils.RespondJSON(c, 200, rs, alert)
	// }
}

// Emailverify godoc
// @summary List Auths
// @description List all auths
// @tags auths
// @security ApiKeyAuth
// @id Emailverify
// @accept json
// @produce json
// @response 200 {array} []models.Login "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /emailverify [get]
func Emailverify(c *gin.Context) {
	// var rs []models.Login
	// err := services.GetAllAuth(&rs)
	// alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"}
	// if err != nil {
	// 	alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"}
	// 	utils.RespondJSON(c, 404, []string{}, alert)
	// } else {
	// 	utils.RespondJSON(c, 200, rs, alert)
	// }
}
