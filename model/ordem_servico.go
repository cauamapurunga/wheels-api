package model

import "time"

type OrdemServico struct {
	Id               int       `json:"id"`
	DescricaoServico string    `json:"descricao_servico"`
	Custo            float64   `json:"custo"`
	DataServico      time.Time `json:"data_servico"`
	VeiculoPlaca     string    `json:"veiculo_placa"`
}
