package app

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

func StarOnMe(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, nil)
	req.Header.Set("Authorization", "token 50c93442b8b519adfda8586a8f6a9418f4fac967")
	req.Header.Set("Content-Length", "0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", resp.Status)
}

func RemoveStarOnMe(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", "token 50c93442b8b519adfda8586a8f6a9418f4fac967")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", resp.Status)
}

func BasicAuth(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token 50c93442b8b519adfda8586a8f6a9418f4fac967")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	arr := make([]repoInformation,0)
	json.Unmarshal(bodyText, &arr)

	for key, value := range arr {
		log.Printf("[%v] -> %v\n", key, value)
	}

	return s
}

type owner struct {
	Login string `json:"login"`
}
type repoInformation struct {
	Owner owner `json:"owner"`
	Language string `json:"language"`
	FullName string `json:"full_name"`
	StargazersCount int `json:"stargazers_count"`
}