package repo

import (
	"database/sql"
	"fmt"
	"laundry/model"
)

type EmployeRepo interface {
	GetEmployeeById(int) (*model.EmployeeModel, error)
	InsertEmployee(*model.EmployeeModel) error
	UpdateEmployee(int, *model.EmployeeModel) error
	DeleteEmployee(int) error
	GetAllEmployee() ([]model.EmployeeModel, error)
}
type employeRepoImpl struct {
	db *sql.DB
}

func (employeRepo *employeRepoImpl) GetEmployeeById(id int) (*model.EmployeeModel, error) {
	qry := "SELECT id, employ_name FROM ms_employee WHERE id = $1"

	employ := &model.EmployeeModel{}
	err := employeRepo.db.QueryRow(qry, id).Scan(&employ.Id, &employ.NamaEmployee)
	if err != nil {
		if err != sql.ErrNoRows{
			return nil, nil
		}
		return nil, fmt.Errorf("error on *employeRepoImpl.GetEmployeeByI : %w", err)
	}
	return employ, nil
}

func (employeRepo *employeRepoImpl) InsertEmployee(employ *model.EmployeeModel) error {
	qry := "INSERT INTO ms_employee(employ_name) VALUES ($1)"
 
	_, err := employeRepo.db.Exec(qry, employ.NamaEmployee)
	if err != nil {
		return fmt.Errorf("error on employeRepoImpl.InsertEmployee : %w", err)
	}
	return nil
}

func(employeRepo *employeRepoImpl) UpdateEmployee(id int, employ *model.EmployeeModel) error{
	qryGetEmploytId := "SELECT id FROM ms_employee WHERE id = $1"

	err := employeRepo.db.QueryRow(qryGetEmploytId, employ.Id).Scan(&employ.Id)
	if err != nil{
		return fmt.Errorf("data not found")
	}

	qry := "UPDATE ms_employee SET employ_name=$1 WHERE id = $2"
	_, err = employeRepo.db.Exec(qry, &employ.NamaEmployee, &employ.Id)
	if err != nil {
		return fmt.Errorf("err on EmployRepoImpl.UpdateEmployee : %w ", err)
	}
	return nil
}

func (EmployeRepo *employeRepoImpl) DeleteEmployee(id int) error {
	qry := "DELETE FROM ms_employee WHERE id = $1"
	_, err := EmployeRepo.db.Exec(qry, id)
	if err != nil{
		return fmt.Errorf("error on  employeRepoImpl.DeleteEmployee : %w ", err)
	}
	return nil
}

func (employeRepo *employeRepoImpl) GetAllEmployee() ([]model.EmployeeModel, error){
	qry := "SELECT id, employ_name FROM ms_employee"
	rows, err := employeRepo.db.Query(qry)
	if  err != nil{
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceById : %w", err)
	}
	defer rows.Close()
	var arrEmploy []model.EmployeeModel
	for rows.Next(){
		employ := &model.EmployeeModel{}
		rows.Scan(&employ.Id, &employ.NamaEmployee)
		arrEmploy = append(arrEmploy, *employ)
	}
	return arrEmploy, nil
}

func NewEmployeeRepo(db *sql.DB) EmployeRepo{
	return &employeRepoImpl{
		db: db,
	}
} 