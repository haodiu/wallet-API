package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/models"
	"wallet/services"
)

type DocumentController struct {
	DocumentService services.DocumentService
}

func NewDocumentController(documentService services.DocumentService) DocumentController  {
	return DocumentController{
		DocumentService: documentService,
	}
}

func (dc *DocumentController) CreateDocument(ctx *gin.Context) {
	var document models.Document
	if err := ctx.ShouldBindJSON(&document); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := dc.DocumentService.CreateDocument(&document)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (dc *DocumentController) GetDocuments(ctx *gin.Context) {
	val := ctx.Param("val")
	documents, err := dc.DocumentService.GetDocuments(val)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, documents)
}

func (dc *DocumentController) DeleteDocument(ctx *gin.Context) {
	id := ctx.Param("id")
	err := dc.DocumentService.DeleteDocument(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (dc *DocumentController) RegisterDocumentRoute(router *gin.RouterGroup)  {
	router.POST("/", dc.CreateDocument)
	router.GET("/:val", dc.GetDocuments)
	router.DELETE("/:id", dc.DeleteDocument)
}