package usecase

import (
	"laundry/model"
	"laundry/repo"

)

type ServiceUsecase interface {
	GetAllService() ([]model.ServiceModel, error)
	GetServiceById(int) (*model.ServiceModel, error)
	InsertService(*model.ServiceModel) error
	UpdateService(int, *model.ServiceModel) error
	DeleteService(int) error
}

type serviceUsecaseImpl struct {
	svcRepo repo.ServiceRepo
}

func (ServiceUsecase *serviceUsecaseImpl) GetServiceById(id int) (*model.ServiceModel, error) {
	return ServiceUsecase.svcRepo.GetServiceById(id)
}

func (ServiceUsecase *serviceUsecaseImpl) InsertService(svc *model.ServiceModel) error {
	return ServiceUsecase.svcRepo.InsertService(svc)
}

func (ServiceUsecase *serviceUsecaseImpl) UpdateService(id int,svc *model.ServiceModel) error  {
	return ServiceUsecase.svcRepo.UpdateService(id, svc)
}

func (ServiceUsecase *serviceUsecaseImpl) DeleteService(id int) error  {
	return ServiceUsecase.svcRepo.DeleteService(id)
}

func (ServiceUsecase *serviceUsecaseImpl) GetAllService() ([]model.ServiceModel, error)  {
	return ServiceUsecase.svcRepo.GetAllService()
}
func NewServiceUseCase(svcRepo repo.ServiceRepo) ServiceUsecase{
	return &serviceUsecaseImpl{
		svcRepo: svcRepo,
	}
}