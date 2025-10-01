package usecase

import (
	"wheels-api/model"
	"wheels-api/repository"
)

type VeiculoUsecase struct {
	repository repository.VeiculoRepository
}

func NewVeiculoUseCase(repo repository.VeiculoRepository) VeiculoUsecase {
	return VeiculoUsecase{
		repository: repo,
	}
}

func (v *VeiculoUsecase) GetVeiculos() ([]model.Veiculo, error) {
	return v.repository.GetVeiculos()
}

func (v *VeiculoUsecase) CreateVeiculo(veiculo model.Veiculo) (model.Veiculo, error) {
	veiculoId, err := v.repository.CreateVeiculo(veiculo)
	if err != nil {
		return model.Veiculo{}, err
	}

	veiculo.Id = veiculoId

	return veiculo, nil
}

func (v VeiculoUsecase) GetVeiculoById(id_veiculo int) (*model.Veiculo, error) {
	veiculo, err := v.repository.GetVeiculoById(id_veiculo)
	if err != nil {
		return nil, err
	}
	return veiculo, nil
}

func (v *VeiculoUsecase) UpdateVeiculo(veiculo model.Veiculo) (int64, error) {
	return v.repository.UpdateVeiculo(veiculo)
}

func (v *VeiculoUsecase) DeleteVeiculo(id int) (int64, error) {
	return v.repository.DeleteVeiculo(id)
}

func (v *VeiculoUsecase) PatchVeiculo(id int, fields map[string]interface{}) (int64, error) {
	return v.repository.PatchVeiculo(id, fields)
}
