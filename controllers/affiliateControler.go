package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/models"
	"wallet/services"
)

type AffiliateController struct {
	AffiliateService services.AffiliateService
}

func NewAffiliateController(affiliateService services.AffiliateService) AffiliateController  {
	return AffiliateController{
		AffiliateService: affiliateService,
	}
}

func (ac *AffiliateController) CreateAffiliate(ctx *gin.Context) {
	var affiliate models.Affiliate
	if err := ctx.ShouldBindJSON(&affiliate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ac.AffiliateService.CreateAffiliate(&affiliate)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AffiliateController) GetAllAffiliate(ctx *gin.Context) {
	customers, err := ac.AffiliateService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, customers)
}

func (ac *AffiliateController) UpdateAffiliate(ctx *gin.Context) {
	id := ctx.Param("id")
	var affiliate models.Affiliate
	if err := ctx.ShouldBindJSON(&affiliate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ac.AffiliateService.UpdateAffiliate(id, &affiliate)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AffiliateController) DeleteAffiliate(ctx *gin.Context) {
	id := ctx.Param("id")
	err := ac.AffiliateService.DeleteAffiliate(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AffiliateController) RegisterAffiliateRoute(router *gin.RouterGroup) {
	router.POST("/", ac.CreateAffiliate)
	router.GET("/", ac.GetAllAffiliate)
	router.PUT("/update/:id", ac.UpdateAffiliate)
	router.DELETE("/delete/:id", ac.DeleteAffiliate)
}
