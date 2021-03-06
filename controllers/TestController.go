package controllers 
 
// controllers.SetupTestRoutes(v1) 
import ( 
	"apitest/models" 
	"apitest/services" 
	"apitest/utils" 
 
	"github.com/gin-gonic/gin" 
) 
 
func SetupTestRoutes(rg *gin.RouterGroup) { 
	router := rg.Group("/tests") 
	router.GET("", listTest) 
	router.GET("/last/iddesc", listlastTest) 
	router.POST("", addNewTest) 
	router.GET("/by/:id", getOneTest) 
	router.PUT("", putOneTest) 
	router.DELETE("/:id", deleteTest) 
} 
 
// ListTests godoc 
// @summary List Tests 
// @description List all tests 
// @tags tests 
// @security ApiKeyAuth 
// @id ListTests 
// @accept json 
// @produce json 
// @response 200 {array} []models.Test "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/tests [get] 
func listTest(c *gin.Context) { 
	var rs []models.Test 
	err := services.GetAllTest(&rs) 
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"} 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} else { 
		utils.RespondJSON(c, 200, rs, alert) 
	} 
} 
 
// ListTests godoc 
// @summary List Tests 
// @description List all tests 
// @tags tests 
// @security ApiKeyAuth 
// @id ListlastTests 
// @accept json 
// @produce json 
// @response 200 {array} []models.Test "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/tests [get] 
func listlastTest(c *gin.Context) { 
	var rs []models.Test 
	err := services.GetAllIdDescTest(&rs) 
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"} 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} else { 
		utils.RespondJSON(c, 200, rs, alert) 
	} 
} 
 
// CreateTest godoc 
// @summary Create Test 
// @description Create new test 
// @tags tests 
// @security ApiKeyAuth 
// @id CreateTest 
// @accept json 
// @produce json 
// @param Test body models.TestForCreate true "Test data to be created" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/tests [post] 
func addNewTest(c *gin.Context) { 
	var rs models.Test 
	c.BindJSON(&rs) 
	err := services.AddNewTest(&rs) 
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"} 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} else { 
		utils.RespondJSON(c, 200, rs, alert) 
	} 
} 
 
// GetTest godoc 
// @summary Get Test 
// @description  Get test by id 
// @tags tests 
// @security Basic auth 
// @id GetTest 
// @accept json 
// @produce json 
// @param id path int true "id of test to be gotten" 
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 409 {object} utils.ResponseData "Conflict" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/tests/:id [get] 
func getOneTest(c *gin.Context) { 
	id := c.Params.ByName("id") 
	var rs models.Test 
	err := services.GetOneTest(&rs, id) 
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"} 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} else { 
		utils.RespondJSON(c, 200, rs, alert) 
	} 
} 
 
// UpdateTest godoc 
// @summary Update Test 
// @description Update test by id 
// @tags tests 
// @security ApiKeyAuth 
// @id UpdateTest 
// @accept json 
// @produce json 
// @param id path int true "id of test to be updated" 
// @param Test body models.TestForUpdate true "Test data to be updated" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/tests/:id [patch] 
func putOneTest(c *gin.Context) { 
	var rs models.Test 
	id := c.Params.ByName("id") 
	err := services.GetOneTest(&rs, id) 
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"} 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Success", Type: "success"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} 
	c.BindJSON(&rs) 
	err = services.PutOneTest(&rs, id) 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} else { 
		utils.RespondJSON(c, 200, rs, alert) 
	} 
} 
 
// DeleteTest godoc 
// @summary Delete Test 
// @description Delete test by id 
// @tags tests 
// @security ApiKeyAuth 
// @id DeleteTest 
// @accept json 
// @produce json 
// @param id path int true "id of test to be deleted" 
// @response 200 {object} utils.ResponseData "OK" 
// @response 400 {object} utils.ResponseData "Bad Request" 
// @response 401 {object} utils.ResponseData "Unauthorized" 
// @response 500 {object} utils.ResponseData "Internal Server Error" 
// @Router /api/v1/tests/:id [delete] 
func deleteTest(c *gin.Context) { 
	var rs models.Test 
	id := c.Params.ByName("id") 
	err := services.DeleteTest(&rs, id) 
	alert := utils.Alert{Msg: "สำเร็จ", Title: "Success", Type: "success"} 
	if err != nil { 
		alert = utils.Alert{Msg: err.Error(), Title: "Error", Type: "error"} 
		utils.RespondJSON(c, 404, []string{}, alert) 
	} else { 
		utils.RespondJSON(c, 200, rs, alert) 
	} 
} 
 