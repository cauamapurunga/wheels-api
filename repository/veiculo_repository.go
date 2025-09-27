package repository

import (
	"database/sql"
	"fmt"
	"wheels-api/model"
)

type VeiculoRepository struct {
	connection *sql.DB
}

func NewVeiculoRepository(connection *sql.DB) VeiculoRepository {
	return VeiculoRepository{
		connection: connection,
	}
}

func (pr *VeiculoRepository) GetVeiculos() ([]model.Veiculo, error) {
	query := "SELECT id, placa, marca, modelo, ano_fabricacao, cor, nome_proprietario FROM veiculos"

	rows, err := pr.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	veiculos := []model.Veiculo{}

	for rows.Next() {
		var veiculo model.Veiculo
		err := rows.Scan(
			&veiculo.Id,
			&veiculo.Placa,
			&veiculo.Marca,
			&veiculo.Modelo,
			&veiculo.AnoFabricacao,
			&veiculo.Cor,
			&veiculo.NomeProprietario,
		)
		if err != nil {
			return nil, err
		}
		veiculos = append(veiculos, veiculo)
	}

	return veiculos, nil
}

func (pr *VeiculoRepository) GetVeiculoById(id int) (*model.Veiculo, error) {
	query := `SELECT id, placa, marca, modelo, ano_fabricacao, cor, nome_proprietario
			  FROM veiculos WHERE id = $1`

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var veiculo model.Veiculo

	err = stmt.QueryRow(id).Scan(
		&veiculo.Id,
		&veiculo.Placa,
		&veiculo.Marca,
		&veiculo.Modelo,
		&veiculo.AnoFabricacao,
		&veiculo.Cor,
		&veiculo.NomeProprietario,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &veiculo, nil
}

func (pr *VeiculoRepository) CreateVeiculo(veiculo model.Veiculo) (int, error) {
	var id int

	query := `INSERT INTO veiculos (placa, marca, modelo, ano_fabricacao, cor, nome_proprietario)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = stmt.QueryRow(
		veiculo.Placa,
		veiculo.Marca,
		veiculo.Modelo,
		veiculo.AnoFabricacao,
		veiculo.Cor,
		veiculo.NomeProprietario,
	).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	stmt.Close()

	return id, nil
}
