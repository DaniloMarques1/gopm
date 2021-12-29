package model

type Password struct {
	Id   	 string
	Name 	 string
	Pwd  	 string
	MasterId string
}

type PasswordRepository interface {
	Save(*Password) error
	FindByName(string, string) (*Password, error)
	RemoveByName(string) error
	FindAll() ([]Password, error)
}
