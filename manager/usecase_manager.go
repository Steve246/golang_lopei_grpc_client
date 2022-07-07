package manager

import "golang_lopei_grc_client/usecase"

type UseCaseManager interface {
	CheckBalanceUsecase() usecase.CheckBalanceUsecase
}


type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CheckBalanceUsecase() usecase.CheckBalanceUsecase {
	return usecase.NewCheckBalanceUseCase(u.repoManager.CustomerRepo())

}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{repoManager: repoManager}
}