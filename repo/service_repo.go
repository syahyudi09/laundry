package repo

import (
	"database/sql"
	"fmt"
	"laundry/model"

)

type ServiceRepo interface {
	GetAllService() ([]model.ServiceModel, error)
	GetServiceById(int) (*model.ServiceModel, error)
	InsertService(*model.ServiceModel) error
	UpdateService(int, *model.ServiceModel) error
	DeleteService(int) error
}

type serviceRepoImpl struct {
	db *sql.DB
}


func (svceRepo *serviceRepoImpl) GetAllService() ([]model.ServiceModel, error){
	qry := "SELECT id, name, price, uom FROM ms_service ORDER BY id"
	rows, err := svceRepo.db.Query(qry)
	if  err != nil{
		return nil, fmt.Errorf("error on serviceRepoImpl.GetServiceById : %w", err)
	}
	defer rows.Close()
	var arrSvc []model.ServiceModel
	for rows.Next(){
		svc := &model.ServiceModel{}
		rows.Scan(&svc.Id, &svc.Name, &svc.Price, &svc.Uom)
		arrSvc = append(arrSvc, *svc)
	}
	return arrSvc, nil
}

func (svceRepo *serviceRepoImpl) InsertService(svc *model.ServiceModel) error {
	qry := "INSERT INTO ms_service(name, price, uom) VALUES ($1,$2,$3)"
 
	_, err := svceRepo.db.Exec(qry, svc.Name, svc.Price, svc.Uom)
	if err != nil {
		return fmt.Errorf("error on serviceRepoImpl. InsertService : %w", err)
	}
	return nil
}

func (svceRepo *serviceRepoImpl) UpdateService(id int, svc *model.ServiceModel ) error{
	qryGetId := "SELECT id FROM ms_service WHERE id = $1"

	err := svceRepo.db.QueryRow(qryGetId, svc.Id).Scan(&svc.Id)
	if err != nil{
		return fmt.Errorf("data not found")
	}

	qry := "UPDATE ms_service SET name=$1, price=$2, uom=$3 WHERE id = $4"
	_, err = svceRepo.db.Exec(qry, &svc.Name, &svc.Price, &svc.Uom, &svc.Id,)
	if err != nil {
		return fmt.Errorf("err on serviceRepoImpl.UpdateService : %w ", err)
	}
	return nil
}

func (svceRepo *serviceRepoImpl) DeleteService(id int) error{
	qry := "DELETE FROM ms_service WHERE id = $1"
	_, err := svceRepo.db.Exec(qry, id)
	if err != nil{
		return fmt.Errorf("error on serviceRepoImpl.DeleteService : %w ", err)
	}
	return nil
}

func NewServiceRepo(db *sql.DB) ServiceRepo{
	return &serviceRepoImpl{
		db: db,
	}
}