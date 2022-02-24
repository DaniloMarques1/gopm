package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/danilomarques1/gopm/cmd/dto"
	"github.com/danilomarques1/gopm/cmd/util"
)

//const BASE_URL = "https://gopmserver.herokuapp.com"
const BASE_URL = "http://127.0.0.1:8080"

type MasterService struct {
	client *http.Client
}

func NewMasterService() *MasterService {
	client := &http.Client{}
	return &MasterService{client: client}
}

func (ms *MasterService) Register(masterDto dto.MasterRegisterDto) error {
	b, err := json.Marshal(masterDto)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, BASE_URL+"/master", bytes.NewReader(b))
	if err != nil {
		return err
	}
	response, err := ms.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return util.HandleError(response.Body)
	}
	return nil
}

func (ms *MasterService) Access(sessionDto dto.SessionRequestDto) (*dto.SessionResponseDto, error) {
	b, err := json.Marshal(sessionDto)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, BASE_URL+"/session", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	response, err := ms.client.Do(request)
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
	var body dto.SessionResponseDto
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return nil, err
	}

	return &body, nil
}
