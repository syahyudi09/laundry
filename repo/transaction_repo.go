package repo

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"time"
)

type TransacationRepo interface {
	GetTransactionById(int) (*model.TransactionModel, error)
	InsertTransaction(*model.TransactionModel) error
	GetAllTransactions() ([]*model.TransactionModel, error)
}

type transactionRepoImpl struct {
	db *sql.DB
}

func (transactionRepo *transactionRepoImpl) InsertTransaction(trans *model.TransactionModel) error {
	tx, err := transactionRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("InsertTransaction B	egin: %w", err)
	} 

	qry := "INSERT INTO tr_header(start_date, end_date, employee_id, cust_name, phone_no) VALUES ($1, $2, $3, $4, $5) RETURNING no"
	err = tx.QueryRow(qry, trans.StartDate, trans.EndDate, trans.EmployeeId, trans.CustName, trans.Phone).Scan(&trans.No)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on transactionRepoImpl.InsertTransaction Header: %w", err)
	}

	trans.StartDate = time.Now()
	err = transactionRepo.InsertDetailTransaction(tx, trans)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on transactionRepoImpl.InsertTransaction Commit: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error on transactionRepoImpl.InsertTransaction Commit: %w", err)
	}
	return nil
}

func (transactionRepo *transactionRepoImpl) InsertDetailTransaction(tx *sql.Tx, trans *model.TransactionModel) error {
	qry := "INSERT INTO tr_detail(trx_no, service_id, service_name, qty, uom, price) VALUES ($1, $2, $3, $4, $5, $6)"
	for _, det := range trans.ArrDetail {
		var svcRepo ServiceRepo
		service, err := svcRepo.GetServiceById(det.ServiceId)
		if err != nil {
			return fmt.Errorf("error on transactionRepoImpl.InsertTransaction GetServiceById: %w", err)
		}
		det.ServiceName = service.Name
		det.Uom = service.Uom
		det.Price = service.Price

		_, err = tx.Exec(qry, trans.No, service.Id, service.Name, det.Qty, service.Uom, service.Price)
		if err != nil {
			return fmt.Errorf("error on transactionRepoImpl.InsertTransaction Detail: %w", err)
		}
	}
	return nil
}

func (transactionRepo *transactionRepoImpl) GetTransactionById(no int) (*model.TransactionModel, error) {
	qry := "SELECT th.no, th.start_date, th.end_date, th.employee_id, th.cust_name, th.phone_no, td.id, td.trx_no, td.service_id, td.service_name, td.qty, td.uom, td.price FROM tr_header AS th JOIN tr_detail AS td ON th.no = td.trx_no WHERE th.no = $1"

	trans := &model.TransactionModel{}

	rows, err := transactionRepo.db.Query(qry, no)
	if err != nil {
		return nil, fmt.Errorf("error on transactionRepoImpl.GetTransactionById: %w", err)
	}
	defer rows.Close()

	for rows.Next() {	
		detail := model.DetailTransaction{}
		err := rows.Scan(
			&trans.No, &trans.StartDate, &trans.EndDate, &trans.EmployeeId, &trans.CustName, &trans.Phone,
			&detail.Id, &detail.TrxNo, &detail.ServiceId, &detail.ServiceName, &detail.Qty, &detail.Uom, &detail.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("error on transactionRepoImpl.GetTransactionById: %w", err)
		}
		trans.ArrDetail = append(trans.ArrDetail, detail)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error on transactionRepoImpl.GetTransactionById: %w", err)
	}
	return trans, nil
}

func (transactionRepo *transactionRepoImpl) GetAllTransactions() ([]*model.TransactionModel, error) {
	qry := "SELECT th.no, th.start_date, th.end_date, th.employee_id, th.cust_name, th.phone_no, td.id, td.trx_no, td.service_id, td.service_name, td.qty, td.uom, td.price FROM tr_header AS th JOIN tr_detail AS td ON th.no = td.trx_no ORDER BY th.no"

	rows, err := transactionRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on transactionRepoImpl.GetAllTransactions: %w", err)
	}
	defer rows.Close()

	// slice untuk menyimpan semua transaksi yang akan diambil didatabase
	var transactions []*model.TransactionModel
	// untuk melakuakn pemindahian nilai nilai dari
	for rows.Next() {
		trans := &model.TransactionModel{}
		detail := &model.DetailTransaction{}
		err := rows.Scan(
			&trans.No, &trans.StartDate, &trans.EndDate, &trans.EmployeeId, &trans.CustName, &trans.Phone,
			&detail.Id, &detail.TrxNo, &detail.ServiceId, &detail.ServiceName, &detail.Qty, &detail.Uom, &detail.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("error on transactionRepoImpl.GetAllTransactions: %w", err)
		}
		trans.ArrDetail = append(trans.ArrDetail, *detail)
		transactions = append(transactions, trans)
	}
	return transactions, nil
}

func NewTransactionRepo(db *sql.DB) TransacationRepo {
	return &transactionRepoImpl{
		db: db,
	}
}
