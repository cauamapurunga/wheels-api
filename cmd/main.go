package main

import (
	"log"
	"wheels-api/controller"
	"wheels-api/db"
	"wheels-api/repository"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	// Injeção de dependência para Veículos
	veiculoRepository := repository.NewVeiculoRepository(dbConnection)
	veiculoUseCase := usecase.NewVeiculoUseCase(veiculoRepository)
	veiculoController := controller.NewVeiculoController(veiculoUseCase)

	// Injeção de dependência para Ordens de Serviço
	ordemServicoRepository := repository.NewOrdemServicoRepository(dbConnection)
	ordemServicoUseCase := usecase.NewOrdemServicoUseCase(ordemServicoRepository)
	ordemServicoController := controller.NewOrdemServicoController(ordemServicoUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message:": "pong",
			"message":  "pong",
		})
	})

	// Rotas de Veículos
	server.GET("/veiculos", veiculoController.GetVeiculos)
	server.GET("/veiculo/:veiculoId", veiculoController.GetVeiculoById)
	server.POST("/veiculo", veiculoController.CreateVeiculo)
	server.PUT("/veiculo/:veiculoId", veiculoController.UpdateVeiculo)
	server.DELETE("/veiculo/:veiculoId", veiculoController.DeleteVeiculo)

	// Rotas de Ordens de Serviço
	server.POST("/servicos", ordemServicoController.CreateOrdemServico)
	server.GET("/servicos/:veiculoPlaca", ordemServicoController.GetOrdensServicoByPlaca)
	server.PUT("/servicos/:servicoId", ordemServicoController.UpdateOrdemServico)
	server.DELETE("/servicos/:servicoId", ordemServicoController.DeleteOrdemServico)

	server.Run(":8000")
}
