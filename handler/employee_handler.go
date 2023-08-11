package handler

import (
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
}

type employeeHandlerImpl struct {
	employUsecase usecase.EmployeeUsecase
}

func (employHandler *employeeHandlerImpl) GetEmployeeById(ctx *gin.Context){
	idText := ctx.Param("id")
	if idText == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	employ, err := employHandler.employUsecase.GetEmployeeById(id)
	if err != nil{
		fmt.Printf("serviceHandler.GetServiceById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data employee",
		})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{
		"success": true,
		"data": employ,
	})
}

func (employHandler *employeeHandlerImpl) InsertEmployee(ctx *gin.Context){
	employ := &model.EmployeeModel{}
	err := ctx.ShouldBindJSON(&employ)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if employ.NamaEmployee == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	err = employHandler.employUsecase.InsertEmployee(employ)
	if err != nil{
		fmt.Printf("custHandler.custUsecase.InsertCust : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data employee",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (employHandler *employeeHandlerImpl) UpdateEmployee(ctx *gin.Context){
	employ := &model.EmployeeModel{}
	err := ctx.ShouldBindJSON(&employ)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if employ.NamaEmployee == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	err = employHandler.employUsecase.UpdateEmployee(employ.Id, employ)
	if err != nil{
		fmt.Printf(" employHandler.employUsecase.UpdateEmployee : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (employHandler *employeeHandlerImpl) DeleteEmployee(ctx *gin.Context){
	idText := ctx.Param("id")
		if idText == ""{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errorMessage": "Id tidak boleh kosong",
			})
			return
		}

		id, err := strconv.Atoi(idText)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"errorMessage": "Id harus angka",
			})
			return
		}

		err = employHandler.employUsecase.DeleteEmployee(id)
		if err != nil{
			fmt.Printf("employHandler.employUsecase : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika mengambil data service",
			})
			return
		}

		ctx.JSON(http.StatusOK,gin.H{
			"success": true,
		})
}

func (employHandler *employeeHandlerImpl) GetAllEmployee(ctx *gin.Context){
	arrEmploy, err := employHandler.employUsecase.GetAllEmployee()
	if err != nil{
		fmt.Printf("serviceHandler.GetServiceById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"success": true,
		"data": arrEmploy,
	})
}

func NewEmployeeHandler(srv *gin.Engine, employUsecase usecase.EmployeeUsecase) EmployeeHandler{
	employHandler := &employeeHandlerImpl{
		employUsecase: employUsecase,
	}
	srv.GET("/employee/:id", employHandler.GetEmployeeById)
	srv.POST("/employee", employHandler.InsertEmployee)
	srv.PUT("/employee", employHandler.UpdateEmployee)
	srv.DELETE("/employee/:id", employHandler.DeleteEmployee)
	srv.GET("/employee", employHandler.GetAllEmployee)
	return employHandler
}