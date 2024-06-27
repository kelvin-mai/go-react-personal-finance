package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/kelvin-mai/personal-finance/internal/model"
)

type CategoryService struct {
	db *sqlx.DB
}

func NewCategoryService(db *sqlx.DB) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

func (s *CategoryService) GetAllByUserId(userId string) ([]model.Category, error) {
	categories := []model.Category{}
	err := s.db.Select(&categories, "select * from category where user_id = $1", userId)
	return categories, err
}

func (s *CategoryService) GetById(id string) (*model.Category, error) {
	category := model.Category{}
	err := s.db.Get(&category, "select * from category where id = $1", id)
	return &category, err
}

func (s *CategoryService) Create(userId, title string) (*model.Category, error) {
	rows, err := s.db.Queryx(
		`insert into category (user_id, title)
	 values ($1, $2)
	 returning *`,
		userId,
		title,
	)
	if err != nil {
		return nil, err
	}
	category := model.Category{}
	rows.Next()
	err = rows.StructScan(&category)
	return &category, err
}

func (s *CategoryService) Update(id, userId, title string) (*model.Category, error) {
	rows, err := s.db.Queryx(
		`update category
			  set title = $2
			where id = $1
			  and user_id = $3
	   returning *`,
		id,
		title,
		userId,
	)
	if err != nil {
		return nil, err
	}
	category := model.Category{}
	rows.Next()
	err = rows.StructScan(&category)
	return &category, err
}

func (s *CategoryService) Delete(id, userId string) (*model.Category, error) {
	rows, err := s.db.Queryx(
		`delete from category
		 where id = $1
		   and user_id = $2
	   returning *`,
		id,
		userId,
	)
	if err != nil {
		return nil, err
	}
	category := model.Category{}
	rows.Next()
	err = rows.StructScan(&category)
	return &category, err
}
