package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/danilomarques1/gopm/cmd/dto"
	"github.com/danilomarques1/gopm/cmd/util"
)

type PasswordService struct {
	Token  string
	client *http.Client 
}

func NewPasswordService(token string) *PasswordService {
	client := &http.Client{}
	return &PasswordService{Token: token, client: client}
}

func (ps *PasswordService) GetPassword(key string) (*dto.PasswordResponseDto, error) {
	request, err := http.NewRequest(http.MethodGet, BASE_URL+"/password/"+key, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+ps.Token)
	response, err := ps.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, util.HandleError(response.Body)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var body dto.PasswordResponseDto
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return nil, err
	}
	return &body, nil
}
