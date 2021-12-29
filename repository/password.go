package repository

import (
	"database/sql"

	"github.com/danilomarques1/gopm/model"
)

type PasswordRepositoryImpl struct {
	db *sql.DB
}

func NewPasswordRepository(db *sql.DB) *PasswordRepositoryImpl {
	return &PasswordRepositoryImpl{db: db}
}

func (pr *PasswordRepositoryImpl) Save(password *model.Password) error {
	return nil
}

func (pr *PasswordRepositoryImpl) FindByName(name string) (*model.Password, error) {
	return nil, nil
}

func (pr *PasswordRepositoryImpl) RemoveByName(name string) error {
	return nil
}
