package handler

import (
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceHandler interface {
}

type serviceHandlerImpl struct {
		svcUsecase usecase.ServiceUsecase
}

func (svcHandler *serviceHandlerImpl) GetServiceById(ctx *gin.Context){
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

		svc, err := svcHandler.svcUsecase.GetServiceById(id)
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
			"data": svc,
		})
}

func (svcHandler *serviceHandlerImpl) GetAllService(ctx *gin.Context){
	arrSvc, err := svcHandler.svcUsecase.GetAllService()
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
		"data": arrSvc,
	})
}

func (svcHandler *serviceHandlerImpl) InsertService(ctx *gin.Context){
	svc := &model.ServiceModel{}
	err := ctx.ShouldBindJSON(&svc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if svc.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	if svc.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
	}
	
	err = svcHandler.svcUsecase.InsertService(svc)
	if err != nil{
		fmt.Printf("svcHandler.svcUsecase.InsertService : %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data service",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (svcHandler *serviceHandlerImpl) UpdateService(ctx *gin.Context){
	svc := &model.ServiceModel{}
	err := ctx.ShouldBindJSON(&svc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if svc.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Nama Tidak Boleh Kosong",
		})
		return
	}

	if svc.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"errorMessage": "Price harus lebih dari 0",
		})
		return
	}
	
	err = svcHandler.svcUsecase.UpdateService(svc.Id, svc)
	if err != nil{
		fmt.Printf("svcHandler.svcUsecase.InsertService : %v", err.Error())
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

func (svcHandler *serviceHandlerImpl) DeleteService(ctx *gin.Context){
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

		err = svcHandler.svcUsecase.DeleteService(id)
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

func NewServiceHandler(srv *gin.Engine, svcUsecase usecase.ServiceUsecase) ServiceHandler {
		svcHandler := &serviceHandlerImpl{
			svcUsecase: svcUsecase,
		}
		srv.GET("/service", svcHandler.GetAllService)
		srv.GET("/service/:id", svcHandler.GetServiceById)
		srv.POST("/service", svcHandler.InsertService)
		srv.PUT("/service", svcHandler.UpdateService)
		srv.DELETE("/service/:id", svcHandler.DeleteService)

		return svcHandler
}