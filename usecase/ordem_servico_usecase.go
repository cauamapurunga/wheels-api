package usecase

import (
	"context"
	"wheels-api/model"
	"wheels-api/repository"
)

type OrdemServicoUsecase struct {
    repository *repository.OrdemServicoRepository
}

func NewOrdemServicoUseCase(repo *repository.OrdemServicoRepository) *OrdemServicoUsecase {
    return &OrdemServicoUsecase{
        repository: repo,
    }
}


func (u *OrdemServicoUsecase) ListAllOrdensServico(ctx context.Context) ([]model.OrdemServico, error) {
	return u.repository.ListAll(ctx)
}

func (u *OrdemServicoUsecase) CreateOrdemServico(ordem model.OrdemServico) (model.OrdemServico, error) {
	id, err := u.repository.CreateOrdemServico(ordem)
	if err != nil {
		return model.OrdemServico{}, err
	}
	ordem.Id = id
	return ordem, nil
}

func (u *OrdemServicoUsecase) GetOrdensServicoByPlaca(placa string) ([]model.OrdemServico, error) {
	return u.repository.GetOrdensServicoByPlaca(placa)
}

func (u *OrdemServicoUsecase) GetOrdemServicoById(id int) (*model.OrdemServico, error) {
	return u.repository.GetOrdemServicoById(id)
}

func (u *OrdemServicoUsecase) UpdateOrdemServico(ordem model.OrdemServico) (int64, error) {
	return u.repository.UpdateOrdemServico(ordem)
}

func (u *OrdemServicoUsecase) DeleteOrdemServico(id int) (int64, error) {
	return u.repository.DeleteOrdemServico(id)
}
