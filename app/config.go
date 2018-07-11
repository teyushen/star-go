package app

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type Repo struct {
	Owner string
	RepoName string
}

func SaveRepo(repo Repo) {

	writeToConfig(repo)
}

func SaveRepos(repos []Repo) {

	writeToConfig(repos)
}

func AppendRepo(repo Repo) {

	repos := append(ReadRepos(), repo)
	writeToConfig(repos)
}

func AppendRepos(repos []Repo) {

	combineRepo := append(ReadRepos(), repos...)
	writeToConfig(combineRepo)
}

func ReadRepos() []Repo {

	content, err := ioutil.ReadFile("./.config")
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]Repo, 0)
	json.Unmarshal(content, &repos)

	log.Printf("File contents: %s", content)

	return repos
}

func writeToConfig(str interface{}) {
	b, _ := json.Marshal(str)

	err := ioutil.WriteFile("./.config", b, 0644)

	if err != nil {
		panic(err)
	}
	return
}

func WriteToConfig(str interface{}) {
	b, _ := json.Marshal(str)

	err := ioutil.WriteFile("./.config", b, 0644)

	if err != nil {
		panic(err)
	}
	return
}