package handler

import (
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler interface {
}

// untuk memanggil transUsecase
type transactionhandlerImpl struct {
	transUsecase usecase.TransactionUsecase
}

func (transHandler *transactionhandlerImpl) GetTransactionById(ctx *gin.Context) {
	noText := ctx.Param("no")
	if noText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nomor tidak boleh kosong",
		})
		return
	}

	no, err := strconv.Atoi(noText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "ID harus berupa angka",
		})
		return
	}

	trans, err := transHandler.transUsecase.GetTransactionById(no)
	if err != nil {
		if err.Error() == "transaction with ID %d not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success":      false,
				"errorMessage": "Transaksi dengan ID tersebut tidak ditemukan",
			})
			return
		}
		fmt.Printf("transactionHandler.GetTransactionById(): %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan saat mengambil data transaksi",
		})
		return
	}

	if trans.No == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Transaksi dengan ID tersebut tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    trans,
	})
}

func (transHandler *transactionhandlerImpl) InsertTransaction(ctx *gin.Context) {
	trans := &model.TransactionModel{}
	err := ctx.ShouldBindJSON(&trans)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	if trans.EmployeeId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "EmployeeId harus angka",
		})
		return
	}

	if trans.CustName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Nama Customer Tidak Boleh Kosong",
		})
		return
	}

	err = transHandler.transUsecase.InsertTransaction(trans)
	if err != nil {
		fmt.Printf("transactionHandler.GetTransactionById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (transHandler *transactionhandlerImpl) GetAllTransactions(ctx *gin.Context) {
	trans, err := transHandler.transUsecase.GetAllTransactions()
	if err != nil {
		fmt.Printf("transactionHandler.GetTransactionById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    trans,
		"success": true,
	})
}

func NewTransactionHandler(srv *gin.Engine, transUsecase usecase.TransactionUsecase) TransactionHandler {
	transHandler := &transactionhandlerImpl{
		transUsecase: transUsecase,
	}
	srv.GET("/transaction", transHandler.GetAllTransactions)
	srv.GET("/transaction/:no", transHandler.GetTransactionById)
	srv.POST("/transaction", transHandler.InsertTransaction)
	return transHandler
}
