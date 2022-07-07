package manager

import "golang_lopei_grc_client/repository"

type RepositoryManager interface {
	CustomerRepo() repository.CustomerRepository
}

type repositoryManager struct {
	infraManager infraManager
}


func(r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.LopeiClientConn())
}

func NewRepositoryManager(infraManager infraManager) RepositoryManager {
	return &repositoryManager{infraManager: infraManager}
}