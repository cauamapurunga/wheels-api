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

	veiculoRepository := repository.NewVeiculoRepository(dbConnection)
	veiculoUseCase := usecase.NewVeiculoUseCase(veiculoRepository)
	veiculoController := controller.NewVeiculoController(veiculoUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message:": "pong",
		})
	})

	server.GET("/veiculos", veiculoController.GetVeiculos)
	server.GET("/veiculo/:veiculoId", veiculoController.GetVeiculoById)
	server.POST("/veiculo", veiculoController.CreateVeiculo)
	server.PUT("/veiculo/:veiculoId", veiculoController.UpdateVeiculo)
	server.DELETE("/veiculo/:veiculoId", veiculoController.DeleteVeiculo)

	server.Run(":8000")
}
