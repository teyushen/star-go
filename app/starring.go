package app

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)


type User struct {
	Token string
}

func StarOnMe(u User, url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url, nil)
	req.Header.Set("Authorization", "token " + u.Token)
	req.Header.Set("Content-Length", "0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", resp.Status)
}

func RemoveStarOnMe(u User, url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", "token " + u.Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", resp.Status)
}

func CompareStar(u User, url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token " + u.Token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	arr := make([]repoInformation,0)
	json.Unmarshal(bodyText, &arr)

	for key, value := range arr {
		log.Printf("[%v] -> name: %v, star: %v, language: %v", key, value.Name, value.StargazersCount, value.Language)
	}

}

type owner struct {
	Login string `json:"login"`
}

type repoInformation struct {
	Owner owner `json:"owner"`
	Language string `json:"language"`
	Name string `json:name`
	FullName string `json:"full_name"`
	StargazersCount int `json:"stargazers_count"`
}