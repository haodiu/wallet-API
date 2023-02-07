package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/auth"
	"wallet/models"
	"wallet/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	result, err := uc.UserService.CheckUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	tokenStr, errGenerate := auth.GenerateToken(result.Username)
	if errGenerate != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": errGenerate.Error()})
		return
	}
	type response struct {
		Token     string `json:"token"`
		ID        int    `json:"id"`
		Username  string `json:"username"`
		IsAdmin   bool   `json:"is_admin"`
		IsAddInfo bool   `json:"is_add_info"`
	}

	ctx.JSON(http.StatusOK, &response{
		Token:     tokenStr,
		ID:        result.ID,
		Username:  result.Username,
		IsAdmin:   result.IsAdmin,
		IsAddInfo: result.IsAddInfo,
	})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := uc.UserService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(router *gin.RouterGroup)  {
	router.POST("/register", uc.RegisterUser)
	router.POST("/login", uc.Login)
	router.GET("/:id", uc.GetUser)
	router.PUT(":id", uc.UpdateUser)
}