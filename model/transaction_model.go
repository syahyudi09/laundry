package model

import (
	"time"
)

type TransactionModel struct {
	No int
	StartDate time.Time
	EndDate time.Time
	EmployeeId int
	CustName string
	Phone string
	ArrDetail []DetailTransaction
}
type DetailTransaction struct{
	Id int
	TrxNo int
	ServiceId int
	ServiceName string
	Qty float64
	Uom string
	Price float64
}

