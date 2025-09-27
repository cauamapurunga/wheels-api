package repository

import (
	"database/sql"
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
		return 0, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		veiculo.Placa,
		veiculo.Marca,
		veiculo.Modelo,
		veiculo.AnoFabricacao,
		veiculo.Cor,
		veiculo.NomeProprietario,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pr *VeiculoRepository) UpdateVeiculo(veiculo model.Veiculo) error {
	query := `UPDATE veiculos 
			  SET placa = $1, marca = $2, modelo = $3, ano_fabricacao = $4, cor = $5, nome_proprietario = $6
			  WHERE id = $7`

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		veiculo.Placa,
		veiculo.Marca,
		veiculo.Modelo,
		veiculo.AnoFabricacao,
		veiculo.Cor,
		veiculo.NomeProprietario,
		veiculo.Id,
	)

	return err
}

func (pr *VeiculoRepository) DeleteVeiculo(id int) (int64, error) {
	query := `DELETE FROM veiculos WHERE id = $1`

	result, err := pr.connection.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
