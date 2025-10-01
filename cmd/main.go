package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	"wheels-api/controller"
	"wheels-api/db"
	"wheels-api/repository"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	var dbConnection *sql.DB
	var err error
	maxRetries := 10
	retryInterval := 5 * time.Second

	for i := 0; i < maxRetries; i++ {
		dbConnection, err = db.ConnectDB()
		if err == nil {
			break // Connection successful
		}
		log.Printf("could not connect to db: %v. Retrying in %v... (%d/%d)", err, retryInterval, i+1, maxRetries)
		time.Sleep(retryInterval)
	}

	if dbConnection == nil {
		log.Fatalf("Falha ao conectar ao banco de dados após %d tentativas. A aplicação será encerrada.", maxRetries)
	}

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
	server.GET("/veiculos/:veiculoId", veiculoController.GetVeiculoById)
	server.POST("/veiculos", veiculoController.CreateVeiculo)
	server.PUT("/veiculos/:veiculoId", veiculoController.UpdateVeiculo)
	server.PATCH("/veiculos/:veiculoId", veiculoController.PatchVeiculo)
	server.DELETE("/veiculos/:veiculoId", veiculoController.DeleteVeiculo)

	// Rotas de Ordens de Serviço
	server.GET("/servicos", ordemServicoController.GetOrdensServico)
	server.POST("/servicos", ordemServicoController.CreateOrdemServico)
	server.GET("/servicos/:veiculoPlaca", ordemServicoController.GetOrdensServicoByPlaca)
	server.PUT("/servicos/:servicoId", ordemServicoController.UpdateOrdemServico)
	server.PATCH("/servicos/:servicoId", ordemServicoController.PatchOrdemServico)
	server.DELETE("/servicos/:servicoId", ordemServicoController.DeleteOrdemServico)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão para ambiente local
	}
	server.Run(":" + port)
}