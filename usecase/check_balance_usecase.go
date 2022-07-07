package usecase

import "golang_lopei_grc_client/repository"

type CheckBalanceUsecase interface {
	GetBalance(lopeiId int32) (float32, error)
}

type checkBalanceUsecase struct {
	repo repository.CustomerRepository
}

func (c *checkBalanceUsecase) GetBalance(lopeiId int32) (float32, error) {
	return c.repo.CheckBalance(lopeiId)
}

func NewCheckBalanceUseCase(repo repository.CustomerRepository) CheckBalanceUsecase {
	return &checkBalanceUsecase{repo:repo}
}