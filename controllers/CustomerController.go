package controllers

import "github.com/gin-gonic/gin"

type CustomerHandler struct {
}

// GetCustomer godoc
// @summary Get Customer
// @description  Get customer by id
// @tags customers
// @security Basic auth
// @id GetCustomer
// @accept json
// @produce json
// @param id path int true "id of customer to be gotten"
// @response 200 {object} models.Customer "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/customers/:id [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	//Logic goes here
}

// ListCustomers godoc
// @summary List Customers
// @description List all customers
// @tags customers
// @security OAuth2
// @id ListCustomers
// @accept json
// @produce json
// @response 200 {object} []models.Customers "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 409 {object} utils.ResponseData "Conflict"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/customers [get]
func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	//Logic goes here
}

// CreateCustomer godoc
// @summary Create Customer
// @description Create new customer
// @tags customers
// @security ApiKeyAuth
// @id CreateCustomer
// @accept json
// @produce json
// @param Customer body models.CustomerForCreate true "Customer data to be created"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	//Logic goes here
}

// DeleteCustomer godoc
// @summary Delete Customer
// @description Delete customer by id
// @tags customers
// @security ApiKeyAuth
// @id DeleteCustomer
// @accept json
// @produce json
// @param id path int true "id of customer to be deleted"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/customers/:id [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	//Logic goes here
}

// UpdateCustomer godoc
// @summary Update Customer
// @description Update customer by id
// @tags customers
// @security ApiKeyAuth
// @id UpdateCustomer
// @accept json
// @produce json
// @param id path int true "id of customer to be updated"
// @param Customer body models.CustomerForUpdate true "Customer data to be updated"
// @response 200 {object} utils.ResponseData "OK"
// @response 400 {object} utils.ResponseData "Bad Request"
// @response 401 {object} utils.ResponseData "Unauthorized"
// @response 500 {object} utils.ResponseData "Internal Server Error"
// @Router /api/v1/customers/:id [patch]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	//Logic goes here
}
