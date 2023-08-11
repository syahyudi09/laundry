package usecase

import (
	"laundry/model"
	"laundry/repo"
)

type EmployeeUsecase interface {
	GetEmployeeById(int) (*model.EmployeeModel, error)
	InsertEmployee(*model.EmployeeModel) error
	UpdateEmployee(int, *model.EmployeeModel) error
	DeleteEmployee(int) error
	GetAllEmployee() ([]model.EmployeeModel, error)
}

type employeeUsecase struct {
	employRepo repo.EmployeRepo
}

func (EmployeeUsecase *employeeUsecase) GetEmployeeById(id int) (*model.EmployeeModel, error) {
	return EmployeeUsecase.employRepo.GetEmployeeById(id)
}

func (EmployeeUsecase *employeeUsecase) InsertEmployee(employ *model.EmployeeModel) error {
	return EmployeeUsecase.employRepo.InsertEmployee(employ)
}

func (EmployeeUsecase *employeeUsecase) UpdateEmployee(id int,employ *model.EmployeeModel) error  {
	return EmployeeUsecase.employRepo.UpdateEmployee(id, employ)
}

func (EmployeeUsecase *employeeUsecase) DeleteEmployee(id int) error  {
	return EmployeeUsecase.employRepo.DeleteEmployee(id)
}

func (EmployeeUsecase *employeeUsecase) GetAllEmployee() ([]model.EmployeeModel, error)  {
	return EmployeeUsecase.employRepo.GetAllEmployee()
}

func NewEmployeeUsecase(employRepo repo.EmployeRepo) EmployeeUsecase {
	return &employeeUsecase{
		employRepo: employRepo,
	}
}
