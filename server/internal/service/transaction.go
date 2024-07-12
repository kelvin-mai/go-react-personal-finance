package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/kelvin-mai/personal-finance/internal/model"
)

type TransactionService struct {
	db *sqlx.DB
}

func NewTransactionService(db *sqlx.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

func (s *TransactionService) GetAll(userId string, page, pageSize int) ([]model.Transaction, error) {
	offset := (page - 1) * pageSize
	transactions := []model.Transaction{}
	err := s.db.Select(&transactions,
		`select * from transactions
		 where user_id = $1
		 limit $2 offset $3`,
		userId,
		pageSize,
		offset,
	)
	return transactions, err
}

func (s *TransactionService) GetById(id string) (*model.Transaction, error) {
	transaction := model.Transaction{}
	err := s.db.Get(&transaction, "select * from transactions where id = $1", id)
	return &transaction, err
}

func (s *TransactionService) Create(userId string, input model.TransactionInput) (*model.Transaction, error) {
	prepared := model.Transaction{
		UserId:     userId,
		CategoryId: input.CategoryId,
		Title:      input.Title,
		Amount:     input.Amount,
		Currency:   input.Currency,
		Type:       input.Type,
	}
	rows, err := s.db.NamedQuery(
		`insert into transactions
		(user_id, category_id, title, amount, currency, type)
		values (:user_id, :category_id, :title, :amount, :currency, :type)
		returning *`,
		prepared,
	)
	if err != nil {
		return nil, err
	}
	transaction := model.Transaction{}
	rows.Next()
	err = rows.StructScan(&transaction)
	return &transaction, err
}

func (s *TransactionService) Update(id, userId string, input model.TransactionInput) (*model.Transaction, error) {
	prepared := model.Transaction{
		Id:         id,
		UserId:     userId,
		CategoryId: input.CategoryId,
		Title:      input.Title,
		Amount:     input.Amount,
		Currency:   input.Currency,
		Type:       input.Type,
	}
	rows, err := s.db.NamedQuery(
		`update transactions
		    set category_id = :category_id,
				    title = :title,
						amount = :amount,
						currency = :currency,
						type = :type
			where id = :id
			  and user_id = :user_id
		returning *`,
		prepared,
	)
	if err != nil {
		return nil, err
	}
	transaction := model.Transaction{}
	rows.Next()
	err = rows.StructScan(&transaction)
	return &transaction, err
}

func (s *TransactionService) Delete(id, userId string) (*model.Transaction, error) {
	rows, err := s.db.Queryx(
		`delete from transactions
		 where id = $1
		   and user_id = $2
	   returning *`,
		id,
		userId,
	)
	if err != nil {
		return nil, err
	}
	transaction := model.Transaction{}
	rows.Next()
	err = rows.StructScan(&transaction)
	return &transaction, err
}
