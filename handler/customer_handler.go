package handler

import (
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
}

type customerHandlerImpl struct {
		CustomerUsecase usecase.CustomerUsecase
}

func (custHandler *customerHandlerImpl) GetCustomerById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	cust, err := custHandler.CustomerUsecase.GetCustomerById(id)
	if err != nil {
		fmt.Printf("customereHandler.GetCustomerById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cust,
	})
}

func (custHandler *customerHandlerImpl) InsertCustomer(ctx *gin.Context){
	cust := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cust)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if cust.CustName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	if cust.Phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Phone Tidak Boleh Kosong",
		})
	}
	
	err = custHandler.CustomerUsecase.InsertCustomer(cust)
	if err != nil{
		fmt.Printf("custHandler.custUsecase.InsertCust : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data service",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (custHandler *customerHandlerImpl) UpdateCustomer(ctx *gin.Context){
	cust := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cust)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if cust.CustName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	if cust.Phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Phone Tidak Boleh Kosong",
		})
		return
	}
	
	err = custHandler.CustomerUsecase.UpdateCustomer(cust.CustId, cust)
	if err != nil{
		fmt.Printf("custHandler.custUsecase.InsertCust : %v", err.Error())
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

func (custHandler *customerHandlerImpl) DeleteCustomer(ctx *gin.Context){
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

		err = custHandler.CustomerUsecase.DeleteCustomer(id)
		if err != nil{
			fmt.Printf("serviceHandler.DeleteService: %v ", err.Error())
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

func (custHandler *customerHandlerImpl) GetAllCustomer(ctx *gin.Context){
	arrCust, err := custHandler.CustomerUsecase.GetAllCustomer()
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
		"data": arrCust,
	})
}

func NewCustomerHandler(srv *gin.Engine, custUsecase usecase.CustomerUsecase) CustomerHandler {
	custHandler := &customerHandlerImpl{
		CustomerUsecase: custUsecase,
	}
	srv.GET("/customer/:id", custHandler.GetCustomerById)
    srv.POST("/customer", custHandler.InsertCustomer)
	srv.PUT("/customer", custHandler.UpdateCustomer)
	srv.DELETE("/customer/:id", custHandler.DeleteCustomer)
	srv.GET("/customer", custHandler.GetAllCustomer)

	return custHandler
}