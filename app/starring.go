package app

import (
	"sort"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/teyushen/star-go/arrays"
)

func GetRepoInfo(u User, url string) []RepoInformation {

	resp := request("GET", u.Token, url)

	bodyText, _ := ioutil.ReadAll(resp.Body)

	reposInfo := make([]RepoInformation, 0)
	json.Unmarshal(bodyText, &reposInfo)

	return reposInfo
}

func CollectAllReposInfo(u User, urls ...string) []RepoInformation {

	reposInfo := make([]RepoInformation, 0)
	ch := make(chan []RepoInformation)

	for _, url := range urls {
		go func() {
			ch <- GetRepoInfo(u, url)
		}()
		reposInfo = append(reposInfo, <- ch...)
	}
	return reposInfo
}

func PrepareReposInfo() []RepoInformation{
	u := GetUser()
	repos := GetRepos()

	arr := make([]string, 0)
	for _, repo := range repos {
		arr = append(arr, "https://api.github.com/users/" + repo.Owner + "/repos")
	}

	reposInfo := CollectAllReposInfo(u, arr...)
	compareReposInfo := make([]RepoInformation, 0)
	for _, repo := range repos {
		for _, repoInfo := range reposInfo {
			if repo.Owner == repoInfo.Owner.Login && arrays.Contains(repo.RepoNames, repoInfo.Name) {
				compareReposInfo = append(compareReposInfo, repoInfo)
			}

		}
	}
	return compareReposInfo
}

type byStar []RepoInformation

func (a byStar) Len() int           { return len(a) }
func (a byStar) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStar) Less(i, j int) bool { return a[i].StargazersCount > a[j].StargazersCount }


func CompareStar(reposInfo []RepoInformation) {

	sort.Sort(byStar(reposInfo))

	//fmt.Println("|Top|---------------|star|---------------|language|--------------------|name|")
	fmt.Printf("%5v %15v %15v %50v \n", "TOP", "STAR", "LANGUAGE", "FULLNAME")
	for index, value := range reposInfo {
		fmt.Printf("%5v %15v %15v %50v \n", index+1, value.StargazersCount, value.Language, value.FullName)
	}

}

type Owner struct {
	Login string `json:"login"`
}

type RepoInformation struct {
	Owner Owner `json:"owner"`
	Language string `json:"language"`
	Name string `json:name`
	FullName string `json:"full_name"`
	StargazersCount int `json:"stargazers_count"`
}