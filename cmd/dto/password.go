package dto

type PasswordRequestDto struct {
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type PasswordResponseDto struct {
	Id  string `json:"id"`
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}

type KeysResponseDto struct {
	Keys []string `json:"keys"`
}
