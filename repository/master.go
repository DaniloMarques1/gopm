package repository

import (
	"database/sql"

	"github.com/danilomarques1/gopm/model"
)

type MasterRepositoryImpl struct {
	db *sql.DB
}

func NewMasterRepository(db *sql.DB) *MasterRepositoryImpl {
	return &MasterRepositoryImpl{
		db: db,
	}
}

func (mr *MasterRepositoryImpl) Save(master *model.Master) error {
	return nil
}

func (mr *MasterRepositoryImpl) FindByPassword(pwd string) (*model.Master, error) {
	return nil, nil
}
