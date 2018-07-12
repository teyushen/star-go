package app

import (
	"net/http"
	"log"
)

func request(method, token string, url string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", "token " + token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}