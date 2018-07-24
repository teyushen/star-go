package app

import (
	"sort"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strings"
	"log"
	"sync"
)

func GetRepoInfo(u User, url string) RepoInformation {

	resp := request("GET", u.Token, url)

	bodyText, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(bodyText), "Bad credentials") {
		fmt.Println(string(bodyText), "Token: ", u.Token)
	}

	reposInfo := RepoInformation{}
	json.Unmarshal(bodyText, &reposInfo)

	return reposInfo
}

func CollectAllReposInfo(u User, urls ...string) []RepoInformation {
	reposInfo := make([]RepoInformation, 0)
	ch := make(chan RepoInformation, len(urls))
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			ch <- GetRepoInfo(u, url)
		}(url)
	}
	wg.Wait()
	close(ch)
	for r := range ch {
		reposInfo = append(reposInfo, r)
	}
	return reposInfo
}

func PrepareReposInfo(number int) []RepoInformation{
	u := GetUser()
	repos := GetRepos(number)

	arr := make([]string, 0)
	for _, repo := range repos {
		//arr = append(arr, "https://api.github.com/users/" + repo.Owner + "/repos")
		//arr = append(arr, "https://api.github.com/orgs/" + repo.Owner + "/repos")
		for _, repoName := range repo.RepoNames {
			arr = append(arr, "https://api.github.com/repos/" + repo.Owner + "/" + repoName)
		}
	}

	log.Println(arr)
	if len(arr) == 0 {
		fmt.Println("Can not find any match repository")
	}
	//
	//reposInfo := CollectAllReposInfo(u, arr...)
	//compareReposInfo := make([]RepoInformation, 0)
	//for _, repo := range repos {
	//	for _, repoInfo := range reposInfo {
	//		//log.Println(repo.Owner, repoInfo.Owner.Login, arrays.Contains(repo.RepoNames, repoInfo.Name))
	//		if repo.Owner == repoInfo.Owner.Login && arrays.Contains(repo.RepoNames, repoInfo.Name) {
	//			compareReposInfo = append(compareReposInfo, repoInfo)
	//		}
	//
	//	}
	//}
	return CollectAllReposInfo(u, arr...)
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