package usecase

import (
	"laundry/model"
	"laundry/repo"
)

type CustomerUsecase interface {
	GetCustomerById(int) (*model.CustomerModel, error)
	InsertCustomer(*model.CustomerModel) error
	UpdateCustomer(int, *model.CustomerModel) error
	DeleteCustomer(int) error
	GetAllCustomer() ([]model.CustomerModel, error)
}

type customerUsecaseImpl struct {
	custRepo repo.CustomerRepo
}

func (CustomerUsecase *customerUsecaseImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	return CustomerUsecase.custRepo.GetCustomerById(id)
}

func (CustomerUsecase *customerUsecaseImpl) InsertCustomer(cust *model.CustomerModel) error {
	return CustomerUsecase.custRepo.InsertCustomer(cust)
}

func (CustomerUsecase *customerUsecaseImpl) UpdateCustomer(id int,svc *model.CustomerModel) error  {
	return CustomerUsecase.custRepo.UpdateCustomer(id, svc)
}

func (CustomerUsecase *customerUsecaseImpl) DeleteCustomer(id int) error  {
	return CustomerUsecase.custRepo.DeleteCustomer(id)
}

func (CustomerUsecase *customerUsecaseImpl) GetAllCustomer() ([]model.CustomerModel, error)  {
	return CustomerUsecase.custRepo.GetAllCustomer()
}

// 
func NewCustomerUsecase(custRepo repo.CustomerRepo) CustomerUsecase {
	return &customerUsecaseImpl{
		custRepo: custRepo,
	}
}