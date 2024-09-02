package web

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func WebRequest(url string, postdata string) ([]byte, int, error) {
	var body []byte
	resp, err := http.Post(url, "application/json", strings.NewReader(postdata))
	if err != nil {
		return body, 0, errors.New("Could not send request" + err.Error())
	}

	defer resp.Body.Close()
	statusCode := resp.StatusCode
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, statusCode, nil
	}

	return body, statusCode, nil
}
