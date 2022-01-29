package dto

type MasterRegisterDto struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type ErrorDto struct {
	Message string `json:"message"`
}
