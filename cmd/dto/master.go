package dto

type MasterRegisterDto struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type ErrorDto struct {
	Message string `json:"message"`
}

type SessionRequestDto struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

type SessionResponseDto struct {
	Token string `json:"token"`
}
