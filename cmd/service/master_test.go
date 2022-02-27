package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/danilomarques1/gopm/cmd/dto"
)

// represents the server errors
const (
	INVALID_TOKEN = `{"message": "Invalid token"}`
	INVALID_BODY  = `{"message": "invalid body"}`
	INVALID_EMAIL = `{"message": "Email already in use"}`
)

type readerMock struct {
	reader io.Reader
}

func NewReaderMock(s string) *readerMock {
	reader := strings.NewReader(s)
	return &readerMock{reader: reader}
}

func (rm *readerMock) Close() error {
	return nil
}

func (rm *readerMock) Read(p []byte) (n int, err error) {
	return rm.reader.Read(p)
}

type DoRequestFunc func(request *http.Request) (*http.Response, error)

func DoRegister(request *http.Request) (*http.Response, error) {
	response := &http.Response{}
	var body dto.MasterRegisterDto
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = NewReaderMock(INVALID_BODY)
		return response, nil
	}

	if len(body.Email) == 0 {
		return nil, errors.New("Invalid body")
	}

	if body.Email == "fitz@mail.com" {
		response.StatusCode = http.StatusBadRequest
		response.Body = NewReaderMock(INVALID_EMAIL)
		return response, nil
	}

	response.StatusCode = http.StatusCreated
	response.Body = NewReaderMock("") // we need it for the close method
	return response, nil
}

func DoAccess(request *http.Request) (*http.Response, error) {
	response := &http.Response{}
	var body dto.SessionRequestDto
	if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Body = NewReaderMock(INVALID_BODY)
		return response, nil
	}

	if len(body.Email) == 0 {
		return nil, errors.New("Invalid body")
	}

	if body.Email == "fitz@mail.com" {
		response.StatusCode = http.StatusBadRequest
		response.Body = NewReaderMock(INVALID_EMAIL)
		return response, nil
	}

	response.StatusCode = http.StatusOK
	responseBody := &dto.SessionResponseDto{Token: "token"}
	b, err := json.Marshal(responseBody)
	if err != nil {
		return nil, err
	}
	response.Body = NewReaderMock(string(b))
	return response, nil
}

// mock the http client
type clientMock struct {
	DoRequest DoRequestFunc
}

func NewClientMock(Do DoRequestFunc) *clientMock {
	return &clientMock{DoRequest: Do}
}

func (client *clientMock) Do(request *http.Request) (*http.Response, error) {
	return client.DoRequest(request)
}

func TestRegister(t *testing.T) {
	cases := []struct {
		label             string
		body              dto.MasterRegisterDto
		shouldReturnError bool
	}{
		{"ShouldRegisterMaster", dto.MasterRegisterDto{Email: "test@mail.com", Pwd: "123456"}, false},
		{"ShouldReturnError", dto.MasterRegisterDto{}, true},
		{"ShouldRegisterMaster", dto.MasterRegisterDto{Email: "fitz@mail.com", Pwd: "123456"}, true},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			masterService := NewMasterService(NewClientMock(DoRegister))
			err := masterService.Register(tc.body)
			returnedErr := err != nil
			if returnedErr != tc.shouldReturnError {
				t.Fatalf("Should return error: %v. instead got: %v\n", tc.shouldReturnError, err)
			}
		})
	}
}

func TestAccess(t *testing.T) {
	cases := []struct {
		label             string
		body              dto.SessionRequestDto
		shouldReturnError bool
		token             string
	}{
		{"ShouldRegisterMaster", dto.SessionRequestDto{Email: "test@mail.com", Pwd: "123456"}, false, "token"},
		{"ShouldReturnError", dto.SessionRequestDto{}, true, ""},
		{"ShouldRegisterMaster", dto.SessionRequestDto{Email: "fitz@mail.com", Pwd: "123456"}, true, ""},
	}

	for _, tc := range cases {
		t.Run(tc.label, func(t *testing.T) {
			masterService := NewMasterService(NewClientMock(DoAccess))
			response, err := masterService.Access(tc.body)
			returnedErr := err != nil
			if returnedErr != tc.shouldReturnError {
				t.Fatalf("Should return error: %v. instead got: %v\n", tc.shouldReturnError, err)
			}

			if response != nil {
				if response.Token != tc.token {
					t.Fatalf("Wrong token returned. Expect: %v got: %v\n", tc.token, response.Token)
				}
			}
		})
	}
}
