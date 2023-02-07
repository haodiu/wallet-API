package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/models"
	"wallet/services"
)

type TransactionController struct {
	TransactionService services.TransactionService
}

func NewTransactionController(transactionService services.TransactionService) TransactionController {
	return TransactionController{
		TransactionService: transactionService,
	}
}

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.TransactionService.CreateTransaction(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TransactionController) GetTransaction(ctx *gin.Context) {
	senderName := ctx.Param("senderName")
	transactions, err := tc.TransactionService.GetTransactions(senderName)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := tc.TransactionService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	err := tc.TransactionService.DeleteTransaction(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (tc *TransactionController) RegisterTransactionRoute(router *gin.RouterGroup)  {
	router.POST("/", tc.CreateTransaction)
	router.GET("/:senderName", tc.GetTransaction)
	router.GET("/", tc.GetAllTransactions)
	router.DELETE("/:id", tc.DeleteTransaction)
}