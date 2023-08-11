package main

import (
	"laundry/handler"
)

func main() {
	handler.NewServer().Run()

	// custRepo := repo.NewCustomerRepo(db)
	// custUseCase := usecase.NewCustomerUsecase(custRepo)
	// handler.NewCustomerHandler(srv ,custUseCase)


	// employRepo := repo.NewEmployeeRepo(db)
	// employUsecase := usecase.EmployeeUsecase(employRepo)
	// handler.NewEmployeeHandler(srv, employUsecase)

}
