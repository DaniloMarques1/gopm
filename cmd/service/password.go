package service

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
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

func (ps *PasswordService) Save(pwdDto *dto.PasswordRequestDto) error {
	b, err := json.Marshal(pwdDto)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, BASE_URL+"/password", bytes.NewReader(b))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", "Bearer "+ps.Token)
	response, err := ps.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		return util.HandleError(response.Body)
	}
	return nil
}

func (ps *PasswordService) RemoveByKey(key string) error {
	request, err := http.NewRequest(http.MethodDelete, BASE_URL+"/password/"+key, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", "Bearer "+ps.Token)
	response, err := ps.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return util.HandleError(response.Body)
	}

	return nil
}

func (ps *PasswordService) Keys() (*dto.KeysResponseDto, error) {
	request, err := http.NewRequest(http.MethodGet, BASE_URL+"/keys", nil)
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
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var keysDto dto.KeysResponseDto
	if err := json.Unmarshal(b, &keysDto); err != nil {
		return nil, err
	}
	return &keysDto, nil
}

func (ps *PasswordService) UpdateByKey(pwdDto *dto.PasswordUpdateRequestDto) error {
	b, err := json.Marshal(pwdDto)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, BASE_URL+"/password", bytes.NewReader(b))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", "Bearer "+ps.Token)
	response, err := ps.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	log.Println(response.StatusCode)

	if response.StatusCode != http.StatusNoContent {
		return util.HandleError(response.Body)
	}

	return nil
}
