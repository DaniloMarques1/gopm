package util

import (
	"io"
	"strings"
	"testing"
)

const (
	ERROR_RESPONSE = `{"message": "Some crazy shit error"}`
)

func mockReader(str string) io.Reader {
	return strings.NewReader(str)
}

func TestHandleError(t *testing.T) {
	cases := []struct {
		label       string
		reader      io.Reader
		message     string
		returnedErr bool
	}{
		{"ShouldParseError", mockReader(ERROR_RESPONSE), "Some crazy shit error", false},
		{"ShouldParseError", mockReader("Wrong json"), "invalid character 'W' looking for beginning of value", true},
	}

	for _, tc := range cases {
		err := HandleError(tc.reader)
		if err.Error() != tc.message {
			t.Fatalf("\nexpected: %v\ngot: %v\n", ERROR_RESPONSE, err)
		}
	}

}
