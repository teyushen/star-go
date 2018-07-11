package main

import (
	"github.com/teyushen/star-go/app"
)

func main() {

	repo := app.Repo{"dennis", "oauth"}
	repo2 := app.Repo{"dennis5", "oauth5"}
	repos := []app.Repo{repo, repo2}

	//app.SaveRepos(repos)
	app.AppendRepos(repos)
	//app.SaveRepo(repo)
	//app.AppendRepo(app.Repo{"dennis2", "oauth2"})

	//repos := make([]app.Repo, 0)
	//r := append(repos, app.Repo{"dennis", "oauth2"}, app.Repo{"dennis1", "oauth2"})

	//log.Print(r)
	//log.Print(repos)


	//f, err := os.Create(".config")
	//f.Close()

	//url := "https://api.github.com/users/teyushen/starred"
	//url1 := "https://api.github.com/user/starred/teyushen/star-go"
	//u := app.User{}
	//u.Token = ""

	//app.StarOnMe(u, url1)
	//
	//app.RemoveStarOnMe(u, url1)
	//app.CompareStar(u, url)
}

