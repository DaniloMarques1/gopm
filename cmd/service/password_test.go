package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/danilomarques1/gopm/cmd/dto"
)

const (
	KEY_ALREADY_IN_USE = `{"message": "Key already in use"}`
	KEY_NOT_FOUND      = `{"message": "Key not found"}`
)

func DoSave(request *http.Request) (*http.Response, error) {
	response := &http.Response{}
	var body dto.PasswordRequestDto
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = NewReaderMock(INVALID_BODY)
		return response, nil
	}

	if body.Key == "github" {
		response.StatusCode = http.StatusBadRequest
		response.Body = NewReaderMock(KEY_ALREADY_IN_USE)
		return response, nil
	}

	response.StatusCode = http.StatusCreated
	responseBody := &dto.PasswordResponseDto{Id: "1", Key: "mail", Pwd: "123456"}
	b, err := json.Marshal(responseBody)
	if err != nil {
		return nil, err
	}

	response.Body = NewReaderMock(string(b))
	return response, nil
}

func DoGetPassword(request *http.Request) (*http.Response, error) {
	response := &http.Response{}

	key := strings.TrimPrefix(request.URL.Path, "/password/")
	if key == "mail" {
		response.StatusCode = http.StatusNotFound
		response.Body = NewReaderMock(KEY_NOT_FOUND)
		return response, nil
	}

	responseBody := &dto.PasswordResponseDto{Id: "1", Key: "mail", Pwd: "123456"}
	b, err := json.Marshal(responseBody)
	if err != nil {
		return nil, err
	}
	response.Body = NewReaderMock(string(b))
	response.StatusCode = http.StatusOK
	return response, nil
}

func TestSavePassword(t *testing.T) {
	cases := []struct {
		label             string
		body              *dto.PasswordRequestDto
		shouldReturnError bool
	}{
		{"ShouldSavePassword", &dto.PasswordRequestDto{Key: "mail", Pwd: "123456"}, false},
		{"ShouldNotSavePassword", &dto.PasswordRequestDto{Key: "github", Pwd: "123456"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			passwordService := NewPasswordService("token", NewClientMock(DoSave)) // TODO fix token
			err := passwordService.Save(tc.body)
			returnedErr := err != nil
			if returnedErr != tc.shouldReturnError {
				t.Fatalf("Should return error: %v. instead got: %v\n", tc.shouldReturnError, err)
			}
		})
	}
}

func TestGetPassword(t *testing.T) {
	cases := []struct {
		label           string
		key             string
		shouldReturnErr bool
	}{
		{"ShouldReturnPassword", "github", false},
		{"ShouldNotReturnPassword", "mail", true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			pwdService := NewPasswordService("token", NewClientMock(DoGetPassword)) // TODO fix token
			_, err := pwdService.GetPassword(tc.key)
			returnedErr := err != nil
			if returnedErr != tc.shouldReturnErr {
				t.Fatalf("Should return error: %v. instead got: %v\n", tc.shouldReturnErr, err)
			}
		})
	}
}
