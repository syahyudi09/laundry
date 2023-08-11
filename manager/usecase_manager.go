package manager

import (
	"laundry/usecase"
	"sync"
)

type UsecaseManager interface {
	GetServiceUsecase() usecase.ServiceUsecase
	GetTransactionUsecase() usecase.TransactionUsecase
}

type usecaseManager struct {
	repoMananger RepoMananger
	svcUsecase usecase.ServiceUsecase
	transUsecase usecase.TransactionUsecase
}

var onceLoadServiceUsecase sync.Once
var onceLoadTransactionUsecase sync.Once

func (um *usecaseManager) GetServiceUsecase() usecase.ServiceUsecase{
	onceLoadServiceUsecase.Do(func() {
		um.svcUsecase = usecase.NewServiceUseCase(um.repoMananger.GetServiceRepo())
	})
	return um.svcUsecase
}

func (um *usecaseManager) GetTransactionUsecase() usecase.TransactionUsecase{
	onceLoadTransactionUsecase.Do(func() {
		um.transUsecase = usecase.NewTransactionUsecase(um.repoMananger.GetTransactionRepo())
	})
	return um.transUsecase
}

func NewUsecaseManager(repoMananger RepoMananger) UsecaseManager {
	return &usecaseManager{
		repoMananger: repoMananger,
	}
}