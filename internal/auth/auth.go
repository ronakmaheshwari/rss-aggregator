package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization");
	if header == "" {
		return "", errors.New("No api key was provided")
	}
	token := strings.Split(header, " ");
	if len(token) != 2 {
		return  "", errors.New("Unauthorized user tried to access this service")
	}

	if token[0] != "Apikey" {
		return  "", errors.New("Unauthorized user tried to access this service")
	}

	return token[1], nil;
}