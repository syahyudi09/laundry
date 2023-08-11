package manager

import (
	"laundry/repo"
	"sync"
)

// gunanya untuk mengunsesiet semua repo

type RepoMananger interface{
	GetServiceRepo() repo.ServiceRepo
	GetTransactionRepo() repo.TransacationRepo
}

type repoMananger struct {
	infraManager InfraManager

	svcRepo repo.ServiceRepo
	transRepo repo.TransacationRepo
}

var onceLoadRepoManager sync.Once

func (rm *repoMananger) GetServiceRepo() repo.ServiceRepo{
	onceLoadRepoManager.Do( func() {
		rm.svcRepo = repo.NewServiceRepo(rm.infraManager.GetDB())
	})
	return rm.svcRepo
}

func (rm *repoMananger) GetTransactionRepo() repo.TransacationRepo{
	onceLoadRepoManager.Do(func() {
		rm.transRepo = repo.NewTransactionRepo(rm.infraManager.GetDB())
	})
	return rm.transRepo
} 

func NewRepoManager(infraManager InfraManager) RepoMananger {
	return &repoMananger{
		infraManager: infraManager,
	}
}

