package usecase

import (
	"laundry/model"
	"laundry/repo"
)

type TransactionUsecase interface {
	GetTransactionById(int) (*model.TransactionModel, error)
	InsertTransaction(*model.TransactionModel) error
	GetAllTransactions() ([]*model.TransactionModel, error) 
}
// untuk memanggil
type transactionUsecasImpl struct {
	transRepo repo.TransacationRepo
}

// function yang menempel ke struct
func (TransactionUsecase *transactionUsecasImpl) GetTransactionById(no int) (*model.TransactionModel, error) {
	return TransactionUsecase.transRepo.GetTransactionById(no)
}

func (TransactionUsecase *transactionUsecasImpl) InsertTransaction(trans *model.TransactionModel) error {
	return TransactionUsecase.transRepo.InsertTransaction(trans)
}
func (TransactionUsecase *transactionUsecasImpl) GetAllTransactions() ([]*model.TransactionModel, error) {
	return TransactionUsecase.transRepo.GetAllTransactions()
}

// function untuk menginisialisasi struct
func NewTransactionUsecase(transRepo repo.TransacationRepo) TransactionUsecase {
	return &transactionUsecasImpl{
		transRepo: transRepo,
	}
}
