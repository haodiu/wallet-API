package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/models"
	"wallet/services"
)

type CustomerController struct {
	CustomerService services.CustomerService
}

func NewCustomerController(customerService services.CustomerService) CustomerController {
	return CustomerController{
		CustomerService: customerService,
	}
}

func (cc *CustomerController) CreateCustomer (ctx *gin.Context) {
	var customer models.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CustomerService.CreateCustomer(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (cc *CustomerController) GetCustomer (ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := cc.CustomerService.GetCustomer(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) GetAll (ctx *gin.Context) {
	customers, err := cc.CustomerService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) UpdateCustomer (ctx *gin.Context) {
	var customer models.Customer
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CustomerService.UpdateCustomer(id, &customer)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (cc *CustomerController) DeleteCustomer (ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.CustomerService.DeleteCustomer(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (cc *CustomerController) RegisterCustomerRoute(router *gin.RouterGroup)  {
	router.POST("/", cc.CreateCustomer)
	router.GET("/:id", cc.GetCustomer)
	router.GET("/", cc.GetAll)
	router.PUT("/:id", cc.UpdateCustomer)
	router.DELETE("/:id", cc.DeleteCustomer)
}