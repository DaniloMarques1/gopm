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
	stmt, err := mr.db.Prepare("insert into master(id, master_pwd) values($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(master.Id, master.MasterPwd)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MasterRepositoryImpl) FindByPassword(pwd string) (*model.Master, error) {
	stmt, err := mr.db.Prepare("select id, master_pwd from master where master_pwd = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var master model.Master
	err = stmt.QueryRow(pwd).Scan(&master.Id, &master.MasterPwd)
	if err != nil {
		return nil, err
	}

	return &master, nil
}
