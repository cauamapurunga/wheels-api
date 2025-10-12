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

func (c *ordemServicoController) GetOrdemServicoById(ctx *gin.Context) {
	idStr := ctx.Param("servicoId")
	servicoId, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID do serviço inválido."})
		return
	}

	ordem, err := c.usecase.GetOrdemServicoById(servicoId)
	if err != nil {
		// Verifica se o erro é "não encontrado" para dar uma resposta mais específica.
		// Assumindo que o usecase retorna sql.ErrNoRows para "não encontrado".
		if err.Error() == "sql: no rows in result set" {
			ctx.JSON(http.StatusNotFound, model.Response{Message: "Ordem de serviço não encontrada."})
			return
		}
		log.Printf("Erro ao buscar ordem de serviço por ID %d: %v", servicoId, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	// O usecase pode retornar nil se não encontrar, mesmo sem erro explícito.
	if ordem == nil {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Ordem de serviço não encontrada."})
		return
	}

	ctx.JSON(http.StatusOK, ordem)
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
	// Verifica se o query parameter "placa" foi passado na URL
	placa := ctx.Query("placa")

	var ordens []model.OrdemServico
	var err error

	if placa != "" {
		// Se a placa foi fornecida, filtra por ela
		ordens, err = c.usecase.GetOrdensServicoByPlaca(placa)
	} else {
		// Caso contrário, lista todas as ordens de serviço
		ordens, err = c.usecase.ListAllOrdensServico(ctx.Request.Context())
	}

	if err != nil {
		log.Printf("Erro ao buscar ordens de serviço (placa: '%s'): %v", placa, err)
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro interno no servidor."})
		return
	}

	if len(ordens) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"servicos": []model.OrdemServico{}})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"servicos": ordens})
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
