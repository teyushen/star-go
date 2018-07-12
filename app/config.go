package app

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type User struct {
	Token string
}

type RepoConfig struct {
	Owner     string
	RepoNames []string
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

	log.Printf("Filename: [.config] -> %s", user)

	return user
}

//func SaveRepo(repo RepoConfig) {
//
//	writeToConfig(".star-go", []RepoConfig{repo})
//}

func SaveRepos(repos []RepoConfig) {

	repoConfigs := make([]RepoConfig, 0)
	for _, repo := range repos {
		if !containRepo(repoConfigs, repo) {
			repoConfigs = append(repoConfigs, repo)
		}
	}
	writeToConfig(".star-go", repoConfigs)
}

func containRepo(repos []RepoConfig, repo RepoConfig) bool {

	for index, value := range repos {
		if value.Owner == repo.Owner {
			repos[index].RepoNames = append(value.RepoNames, repo.RepoNames...)
			return true
		}
	}
	return false

}

func AppendRepos(repos ...RepoConfig) []RepoConfig {

	repoConfigs := GetRepos()
	for _, repo := range repos {
		if !containRepo(repoConfigs, repo) {
			repoConfigs = append(repoConfigs, repo)
		}
	}
	SaveRepos(repoConfigs)
	return repoConfigs
}

func GetRepos() []RepoConfig {

	content, err := ioutil.ReadFile(".star-go")
	if err != nil {
		log.Fatal(err)
	}

	repos := make([]RepoConfig, 0)
	json.Unmarshal(content, &repos)

	log.Printf("Filename: [.star-go] -> %s", content)

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
