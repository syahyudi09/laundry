package repo

import (
	"database/sql"
	"fmt"
	"laundry/model"
)

type CustomerRepo interface {
	GetCustomerById(int) (*model.CustomerModel, error)
	InsertCustomer(*model.CustomerModel) error
	UpdateCustomer(int, *model.CustomerModel) error
	DeleteCustomer(int) error
	GetAllCustomer() ([]model.CustomerModel, error)
	
}

type customerRepoImpl struct {
	db *sql.DB
}

func (custRepo *customerRepoImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	qry := "SELECT cust_id, cust_name, phone FROM ms_customer WHERE cust_id = $1"

	cust := &model.CustomerModel{}
	err := custRepo.db.QueryRow(qry, id).Scan(&cust.CustId, &cust.CustName, &cust.Phone)
	if err != nil {
		if err != sql.ErrNoRows{
			return nil, nil
		}
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceById : %w", err)
	}
	return cust, nil
}

func (custRepo *customerRepoImpl) InsertCustomer(cust *model.CustomerModel) error {
	qry := "INSERT INTO ms_customer(cust_name, phone) VALUES ($1,$2)"
 
	_, err := custRepo.db.Exec(qry, cust.CustName, cust.Phone)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.InsertCustomer. : %w", err)
	}
	return nil
}

func (custRepo *customerRepoImpl) UpdateCustomer(id int, cust *model.CustomerModel ) error{
	qryGetCustId := "SELECT cust_id FROM ms_customer WHERE cust_id = $1"

	err := custRepo.db.QueryRow(qryGetCustId, cust.CustId).Scan(&cust.CustId)
	if err != nil{
		return fmt.Errorf("data not found")
	}

	qry := "UPDATE ms_customer SET cust_name=$1, phone=$2 WHERE cust_id = $3"
	_, err = custRepo.db.Exec(qry, &cust.CustName, &cust.Phone, &cust.CustId)
	if err != nil {
		return fmt.Errorf("err on serviceRepoImpl.UpdateService : %w ", err)
	}
	return nil
}

func (custRepo *customerRepoImpl) DeleteCustomer(id int) error {
	qry := "DELETE FROM ms_customer WHERE cust_id = $1"
	_, err := custRepo.db.Exec(qry, id)
	if err != nil{
		return fmt.Errorf("error on serviceRepoImpl.DeleteService : %w ", err)
	}
	return nil
}

func (custRepo *customerRepoImpl) GetAllCustomer() ([]model.CustomerModel, error){
	qry := "SELECT cust_id, cust_name, phone FROM ms_customer"
	rows, err := custRepo.db.Query(qry)
	if  err != nil{
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceById : %w", err)
	}
	defer rows.Close()
	var arrCust []model.CustomerModel
	for rows.Next(){
		cust := &model.CustomerModel{}
		rows.Scan(&cust.CustId, &cust.CustName, &cust.Phone)
		arrCust = append(arrCust, *cust)
	}
	return arrCust, nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo{
	return &customerRepoImpl{
		db: db,
	}
}