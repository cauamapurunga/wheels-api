package controller

import (
	"log"
	"net/http"
	"strconv"

	"wheels-api/model"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
)

type ordemServicoController struct {
	usecase *usecase.OrdemServicoUsecase
}

func NewOrdemServicoController(u *usecase.OrdemServicoUsecase) *ordemServicoController {
	return &ordemServicoController{
		usecase: u,
	}
}

func (c *ordemServicoController) CreateOrdemServico(ctx *gin.Context) {
	var ordem model.OrdemServico
	if err := ctx.BindJSON(&ordem); err != nil {
		log.Printf("Erro no bind do JSON para criar ordem de serviço: %v", err)
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido."})
		return
	}

	createdOrdem, err := c.usecase.CreateOrdemServico(ordem)
	if err != nil {
		log.Printf("Erro ao criar ordem de serviço: %v", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	ctx.JSON(http.StatusCreated, createdOrdem)
}

func (c *ordemServicoController) GetOrdensServicoByPlaca(ctx *gin.Context) {
	placa := ctx.Param("veiculoPlaca")
	if placa == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "A placa do veículo é obrigatória."})
		return
	}

	ordens, err := c.usecase.GetOrdensServicoByPlaca(placa)
	if err != nil {
		log.Printf("Erro ao buscar ordens de serviço para a placa %s: %v", placa, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if len(ordens) == 0 {
		ctx.JSON(http.StatusOK, []model.OrdemServico{})
		return
	}

	ctx.JSON(http.StatusOK, ordens)
}

func (c *ordemServicoController) UpdateOrdemServico(ctx *gin.Context) {
	idStr := ctx.Param("servicoId")
	servicoId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID do serviço inválido."})
		return
	}

	var ordem model.OrdemServico
	if err := ctx.BindJSON(&ordem); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido."})
		return
	}

	ordem.Id = servicoId

	rowsAffected, err := c.usecase.UpdateOrdemServico(ordem)
	if err != nil {
		log.Printf("Erro ao atualizar ordem de serviço ID %d: %v", servicoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Ordem de serviço não encontrada para atualização."})
		return
	}

	updatedOrdem, err := c.usecase.GetOrdemServicoById(servicoId)
	if err != nil {
		log.Printf("Erro ao buscar ordem de serviço atualizada ID %d: %v", servicoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro ao buscar dados após atualização."})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdem)
}

func (c *ordemServicoController) DeleteOrdemServico(ctx *gin.Context) {
	idStr := ctx.Param("servicoId")
	servicoId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID do serviço inválido."})
		return
	}

	rowsAffected, err := c.usecase.DeleteOrdemServico(servicoId)
	if err != nil {
		log.Printf("Erro ao deletar ordem de serviço ID %d: %v", servicoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Ordem de serviço não encontrada."})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Ordem de serviço deletada com sucesso."})
}

func (c *ordemServicoController) GetOrdensServico(ctx *gin.Context) {
	
	ordens, err := c.usecase.ListAllOrdensServico(ctx.Request.Context())
	if err != nil {
		log.Printf("Erro ao listar ordens de serviço: %v", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if len(ordens) == 0 {
		ctx.JSON(http.StatusOK, []model.OrdemServico{})
		return
	}

	ctx.JSON(http.StatusOK, ordens)
}

func (c *ordemServicoController) PatchOrdemServico(ctx *gin.Context) {
	idStr := ctx.Param("servicoId")
	servicoId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID do serviço inválido."})
		return
	}

	var fields map[string]interface{}
	if err := ctx.BindJSON(&fields); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "Corpo da requisição inválido."})
		return
	}

	rowsAffected, err := c.usecase.PatchOrdemServico(servicoId, fields)
	if err != nil {
		log.Printf("Erro ao atualizar parcialmente ordem de serviço ID %d: %v", servicoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Ordem de serviço não encontrada para atualização."})
		return
	}

	updatedOrdem, err := c.usecase.GetOrdemServicoById(servicoId)
	if err != nil {
		log.Printf("Erro ao buscar ordem de serviço atualizada ID %d: %v", servicoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro ao buscar dados após atualização."})
		return
	}

	ctx.JSON(http.StatusOK, updatedOrdem)
}
