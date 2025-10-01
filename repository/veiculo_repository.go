package repository

import (
	"database/sql"
	"strconv"
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

func (pr *VeiculoRepository) UpdateVeiculo(veiculo model.Veiculo) (int64, error) {
	query := `UPDATE veiculos 
			  SET placa = $1, marca = $2, modelo = $3, ano_fabricacao = $4, cor = $5, nome_proprietario = $6
			  WHERE id = $7`

	stmt, err := pr.connection.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		veiculo.Placa,
		veiculo.Marca,
		veiculo.Modelo,
		veiculo.AnoFabricacao,
		veiculo.Cor,
		veiculo.NomeProprietario,
		veiculo.Id,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (pr *VeiculoRepository) DeleteVeiculo(id int) (int64, error) {
	query := `DELETE FROM veiculos WHERE id = $1`

	result, err := pr.connection.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (pr *VeiculoRepository) PatchVeiculo(id int, fields map[string]interface{}) (int64, error) {
	if len(fields) == 0 {
		return 0, nil
	}

	query := "UPDATE veiculos SET "
	values := []interface{}{}
	i := 1

	for field, value := range fields {
		if i > 1 {
			query += ", "
		}
		switch field {
		case "placa":
			query += "placa = $" + strconv.Itoa(i)
		case "marca":
			query += "marca = $" + strconv.Itoa(i)
		case "modelo":
			query += "modelo = $" + strconv.Itoa(i)
		case "ano_fabricacao":
			query += "ano_fabricacao = $" + strconv.Itoa(i)
		case "cor":
			query += "cor = $" + strconv.Itoa(i)
		case "nome_proprietario":
			query += "nome_proprietario = $" + strconv.Itoa(i)
		}
		values = append(values, value)
		i++
	}
	query += " WHERE id = $" + strconv.Itoa(i)
	values = append(values, id)

	result, err := pr.connection.Exec(query, values...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
