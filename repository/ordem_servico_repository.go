package repository

import (
	"database/sql"
	"wheels-api/model"
)

type OrdemServicoRepository struct {
	connection *sql.DB
}

func NewOrdemServicoRepository(connection *sql.DB) OrdemServicoRepository {
	return OrdemServicoRepository{
		connection: connection,
	}
}

func (r *OrdemServicoRepository) CreateOrdemServico(ordem model.OrdemServico) (int, error) {
	var id int
	query := `INSERT INTO ordens_servico (descricao_servico, custo, data_servico, veiculo_placa)
			  VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.connection.QueryRow(query,
		ordem.DescricaoServico,
		ordem.Custo,
		ordem.DataServico,
		ordem.VeiculoPlaca,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OrdemServicoRepository) GetOrdensServicoByPlaca(placa string) ([]model.OrdemServico, error) {
	query := `SELECT id, descricao_servico, custo, data_servico, veiculo_placa 
			  FROM ordens_servico WHERE veiculo_placa = $1`

	rows, err := r.connection.Query(query, placa)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ordens []model.OrdemServico
	for rows.Next() {
		var ordem model.OrdemServico
		err := rows.Scan(
			&ordem.Id,
			&ordem.DescricaoServico,
			&ordem.Custo,
			&ordem.DataServico,
			&ordem.VeiculoPlaca,
		)
		if err != nil {
			return nil, err
		}
		ordens = append(ordens, ordem)
	}
	return ordens, nil
}

func (r *OrdemServicoRepository) GetOrdemServicoById(id int) (*model.OrdemServico, error) {
	query := `SELECT id, descricao_servico, custo, data_servico, veiculo_placa 
			  FROM ordens_servico WHERE id = $1`

	var ordem model.OrdemServico
	err := r.connection.QueryRow(query, id).Scan(
		&ordem.Id,
		&ordem.DescricaoServico,
		&ordem.Custo,
		&ordem.DataServico,
		&ordem.VeiculoPlaca,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Nenhum erro, apenas não encontrou
		}
		return nil, err
	}

	return &ordem, nil
}

func (r *OrdemServicoRepository) UpdateOrdemServico(ordem model.OrdemServico) (int64, error) {
	query := `UPDATE ordens_servico 
			  SET descricao_servico = $1, custo = $2, data_servico = $3, veiculo_placa = $4
			  WHERE id = $5`

	result, err := r.connection.Exec(query,
		ordem.DescricaoServico,
		ordem.Custo,
		ordem.DataServico,
		ordem.VeiculoPlaca,
		ordem.Id,
	)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *OrdemServicoRepository) DeleteOrdemServico(id int) (int64, error) {
	query := `DELETE FROM ordens_servico WHERE id = $1`

	result, err := r.connection.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
