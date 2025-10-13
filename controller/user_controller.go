package controller

import (
	"net/http"
	"strings"
	"wheels-api/model"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
)

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{useCase: useCase}
}

type UserController struct {
	useCase *usecase.UserUseCase
}

func (uc *UserController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido: " + err.Error()})
		return
	}

	createdUser, err := uc.useCase.Register(user)
	if err != nil {
		if strings.Contains(err.Error(), "E-mail já cadastrado.") {
			c.JSON(http.StatusConflict, model.Response{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, model.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido: " + err.Error()})
		return
	}

	token, err := uc.useCase.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{Message: "Credenciais inválidas."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
