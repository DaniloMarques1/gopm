package model

type Master struct {
	MasterPwd string
}

type MasterRepository interface {
	Save(*Master) error
	FindByPassword(string) (*Master, error)
}
