package app

import (
	"sort"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func GetRepoInfo(u User, url string) []repoInformation {

	resp := request("GET", u.Token, url)

	bodyText, _ := ioutil.ReadAll(resp.Body)

	reposInfo := make([]repoInformation, 0)
	json.Unmarshal(bodyText, &reposInfo)

	return reposInfo
}

func CollectAllReposInfo(u User, urls ...string) []repoInformation {

	reposInfo := make([]repoInformation, 0)
	ch := make(chan []repoInformation)

	for _, url := range urls {
		go func() {
			ch <- GetRepoInfo(u, url)
		}()
		reposInfo = append(reposInfo, <- ch...)
	}

	return reposInfo
}


type byStar []repoInformation

func (a byStar) Len() int           { return len(a) }
func (a byStar) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byStar) Less(i, j int) bool { return a[i].StargazersCount > a[j].StargazersCount }


func CompareStar(reposInfo []repoInformation) {

	sort.Sort(byStar(reposInfo))

	//fmt.Println("|Top|---------------|star|---------------|language|--------------------|name|")
	fmt.Printf("%5v %15v %15v %50v \n", "TOP", "NAME", "LANGUAGE", "FULLNAME")
	for index, value := range reposInfo {
		fmt.Printf("%5v %15v %15v %50v \n", index+1, value.StargazersCount, value.Language, value.FullName)
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