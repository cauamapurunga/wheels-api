package model

type Veiculo struct {
	Id               int    `json:"id"`
	Placa            string `json:"placa"`
	Marca            string `json:"marca"`
	Modelo           string `json:"modelo"`
	AnoFabricacao    int    `json:"ano_fabricacao"`
	Cor              string `json:"cor"`
	NomeProprietario string `json:"nome_proprietario"`
}
