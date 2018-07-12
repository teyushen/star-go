package main

import (
	"os"
	"gopkg.in/urfave/cli.v2"
	"fmt"
	"log"
	"io/ioutil"
)

const StarGoUrl = "https://api.github.com/user/starred/teyushen/star-go"

func init() {
	log.SetOutput(ioutil.Discard)
}


func main() {

	//user := app.GetUser()
	//app.StarOnMe(user, StarGoUrl)
	//app.RemoveStarOnMe(user, StarGoUrl)


	//url := "https://api.github.com/users/teyushen/starred"
	//ur1 := "https://api.github.com/users/believeWang/starred"
	//url1 := "https://api.github.com/user/starred/teyushen/star-go"

	//fmt.Print(app.GetRepoInfo(user, url))
	//fmt.Print(app.CollectAllReposInfo(user, url, ur1))
	//app.CompareStar(app.CollectAllReposInfo(user, url, ur1))


	app := &cli.App{
		Name:        "greet",
		Version:     "0.1.0",
		Description: "This is how we describe greet the app",
		Authors: []*cli.Author{
			{Name: "Harrison", Email: "harrison@lolwut.com"},
			{Name: "Oliver Allen", Email: "oliver@toyshop.com"},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Value: "bob", Usage: "a name to say"},
		},
		Commands: []*cli.Command{
			{
				Name:        "describeit",
				Aliases:     []string{"d"},
				Usage:       "use it to see a description",
				Description: "This is how we describe describeit the function",
				Action: func(c *cli.Context) error {
					fmt.Printf("i like to describe things")
					return nil
				},
			},
		},
	}
	app.Run(os.Args)

	//u := app.User{}
	//u.Token = ""

	//app.RemoveStarOnMe(u, url1)
	//app.CompareStar(user, url)

}

