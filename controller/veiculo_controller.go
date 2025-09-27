package controller

import (
	"log"
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
		log.Printf("Erro ao buscar veículos: %v", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	ctx.JSON(http.StatusOK, veiculos)
}

func (v *veiculoController) GetVeiculoById(ctx *gin.Context) {

	id := ctx.Param("veiculoId")
	if id == "" {
		response := model.Response{
			Message: "Id do veículo não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do veículo precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculo, err := v.veiculoUsecase.GetVeiculoById(veiculoId)
	if err != nil {
		log.Printf("Erro ao buscar veículo por ID %d: %v", veiculoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if veiculo == nil {
		response := model.Response{
			Message: "Veículo não encontrado",
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
		log.Printf("Erro no bind do JSON para criar veículo: %v", err)
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido."})
		return
	}

	insertedVeiculo, err := v.veiculoUsecase.CreateVeiculo(veiculo)

	if err != nil {
		log.Printf("Erro ao criar veículo: %v", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Erro interno no servidor",
		})
		return
	}

	ctx.JSON(http.StatusCreated, insertedVeiculo)
}

func (v *veiculoController) UpdateVeiculo(ctx *gin.Context) {
	id := ctx.Param("veiculoId")
	if id == "" {
		response := model.Response{
			Message: "Id do veículo não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do veículo precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var veiculo model.Veiculo
	if err := ctx.BindJSON(&veiculo); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido."})
		return
	}

	veiculo.Id = veiculoId

	rowsAffected, err := v.veiculoUsecase.UpdateVeiculo(veiculo)
	if err != nil {
		log.Printf("Erro ao atualizar veículo ID %d: %v", veiculoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Veículo não encontrado para atualização."})
		return
	}

	// Busca o veículo atualizado para retornar ao cliente
	updatedVeiculo, err := v.veiculoUsecase.GetVeiculoById(veiculoId)
	if err != nil {
		log.Printf("Erro ao buscar veículo atualizado ID %d: %v", veiculoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro ao buscar dados após atualização."})
		return
	}

	ctx.JSON(http.StatusOK, updatedVeiculo)
}

func (v *veiculoController) DeleteVeiculo(ctx *gin.Context) {
	id := ctx.Param("veiculoId")
	if id == "" {
		response := model.Response{
			Message: "Id do veículo não pode ser nulo.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculoId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do veículo precisa ser um número.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	rowsAffected, err := v.veiculoUsecase.DeleteVeiculo(veiculoId)
	if err != nil {
		log.Printf("Erro ao deletar veículo ID %d: %v", veiculoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if rowsAffected == 0 {
		response := model.Response{
			Message: "Veículo não encontrado.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Veículo deletado com sucesso."})
}
