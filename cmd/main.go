package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	"wheels-api/controller"
	"wheels-api/db"
	"wheels-api/middleware"
	"wheels-api/repository"
	"wheels-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Define que não confiamos em nenhum proxy
	server.SetTrustedProxies(nil)

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

	// User
	userRepository := repository.NewUserRepository(dbConnection)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	veiculoRepository := repository.NewVeiculoRepository(dbConnection)
	veiculoUseCase := usecase.NewVeiculoUseCase(veiculoRepository)
	veiculoController := controller.NewVeiculoController(veiculoUseCase)

	ordemServicoRepository := repository.NewOrdemServicoRepository(dbConnection)
	ordemServicoUseCase := usecase.NewOrdemServicoUseCase(ordemServicoRepository)
	ordemServicoController := controller.NewOrdemServicoController(ordemServicoUseCase)

	server.POST("/register", userController.Register)
	server.POST("/login", userController.Login)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	protected := server.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Rotas de Veículos
	protected.GET("/veiculos", veiculoController.GetVeiculos)
	protected.GET("/veiculos/:veiculoId", veiculoController.GetVeiculoById)
	protected.POST("/veiculos", veiculoController.CreateVeiculo)
	protected.PUT("/veiculos/:veiculoId", veiculoController.UpdateVeiculo)
	protected.PATCH("/veiculos/:veiculoId", veiculoController.PatchVeiculo)
	protected.DELETE("/veiculos/:veiculoId", veiculoController.DeleteVeiculo)

	// Rotas de Ordens de Serviço
	protected.GET("/servicos", ordemServicoController.GetOrdensServico)
	protected.POST("/servicos", ordemServicoController.CreateOrdemServico)
	protected.GET("/servicos/:veiculoPlaca", ordemServicoController.GetOrdensServicoByPlaca)
	protected.PUT("/servicos/:servicoId", ordemServicoController.UpdateOrdemServico)
	protected.PATCH("/servicos/:servicoId", ordemServicoController.PatchOrdemServico)
	protected.DELETE("/servicos/:servicoId", ordemServicoController.DeleteOrdemServico)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Porta padrão para ambiente local
	}
	server.Run(":" + port)
}
