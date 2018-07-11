package app

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type User struct {
	Token string
}

type Repo struct {
	Owner string
	RepoName string
}

func SaveUser(user User) {

	writeToConfig(".config", user)
}

func GetUser() User {

	content, err := ioutil.ReadFile(".config")
	if err != nil {
		log.Fatal(err)
	}

	user := User{}
	json.Unmarshal(content, &user)

	log.Printf("File contents: %s", user)

	return user
}


func SaveRepo(repo Repo) {

	writeToConfig(".star-go", []Repo{repo})
}

func SaveRepos(repos []Repo) {

	writeToConfig(".star-go", repos)
}

func AppendRepo(repo Repo) []Repo{

	repos := append(ReadRepos(), repo)
	SaveRepos(repos)
	return repos
}

func AppendRepos(repos []Repo) {

	combineRepo := append(ReadRepos(), repos...)
	writeToConfig(".star-go", combineRepo)
}

func ReadRepos() []Repo {

	content, err := ioutil.ReadFile(".star-go")
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]Repo, 0)
	json.Unmarshal(content, &repos)

	log.Printf("File contents: %s", content)

	return repos
}

func writeToConfig(filename string, str interface{}) {
	b, _ := json.Marshal(str)

	err := ioutil.WriteFile(filename, b, 0777)

	if err != nil {
		panic(err)
	}
	return
}

