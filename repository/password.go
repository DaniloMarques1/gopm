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
	stmt, err := pr.db.Prepare("select id, name, pwd from password where name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	
	var pwd model.Password
	err = stmt.QueryRow(name).Scan(&pwd.Id, &pwd.Name, &pwd.Pwd)
	if err != nil {
		return nil, err
	}
	return &pwd, nil
}

func (pr *PasswordRepositoryImpl) RemoveByName(name string) error {
	return nil
}

func (pr *PasswordRepositoryImpl) FindAll() ([]model.Password, error) {
	return nil, nil
}
