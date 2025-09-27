package controller

import (
	"net/http"
	"strconv"
	"wheels-api/model"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
)

type veiculoController struct {
	veiculoUsecase usecase.VeiculoUsecase
}

func NewVeiculoController(usecase usecase.VeiculoUsecase) veiculoController {
	return veiculoController{
		veiculoUsecase: usecase,
	}
}

func (v *veiculoController) GetVeiculos(ctx *gin.Context) {
	veiculos, err := v.veiculoUsecase.GetVeiculos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, veiculos)
}

func (v *veiculoController) GetVeiculoById(ctx *gin.Context) {

	id := ctx.Param("veiculoId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculo, err := v.veiculoUsecase.GetVeiculoById(veiculoId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if veiculo == nil {
		response := model.Response{
			Message: "Produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, veiculo)
}

func (v *veiculoController) CreateVeiculo(ctx *gin.Context) {
	var veiculo model.Veiculo
	err := ctx.BindJSON(&veiculo)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedVeiculo, err := v.veiculoUsecase.CreateVeiculo(veiculo)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedVeiculo)
}
