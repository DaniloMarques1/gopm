package util

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/danilomarques1/gopm/cmd/dto"
)

func HandleError(body io.Reader) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	var bodyObj dto.ErrorDto
	if err := json.Unmarshal(bodyBytes, &bodyObj); err != nil {
		return err
	}
	return errors.New(bodyObj.Message)
}
