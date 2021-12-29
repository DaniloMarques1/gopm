package model

type Master struct {
	Id        string
	MasterPwd string
}

type MasterRepository interface {
	Save(*Master) error
	FindByPassword(string) (*Master, error)
}
