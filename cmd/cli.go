package cmd

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
	"github.com/teyushen/star-go/app"
	"os"
)

const StarGoUrl = "https://api.github.com/user/starred/teyushen/star-go"
const StarredUrl = "https://api.github.com/users/teyushen/starred"

func Cli() {
	user := app.GetUser()

	app := &cli.App{
		Name:        "star-go",
		Version:     "0.1.0",
		Usage: 	 "A cli help you easy to show how many stars on Github",
		Description: "This can help us easily compare how many stars of Repository in Github we interested",
		Authors: []*cli.Author{
			{Name: "Dennis", Email: "teyushen@gmail.com"},
		},
		Flags: []cli.Flag{
			//&cli.StringFlag{Name: "name", Value: "bob", Usage: "a name to say"},
		},
		Commands: []*cli.Command{
			{
				Name:        "init",
				Aliases:     []string{},
				ArgsUsage: 	 "<token>",
				//UsageText: 	 "star-go init TOKEN",
				Usage:       "Use it to initial the basic configuration",
				Description: "This is how we add Github token into our config",
				Action: func(c *cli.Context) error {
					if c.Args().Len() == 1 {
						token := c.Args().First()
						fmt.Println("Your Github token is: ", c.Args().First())
						app.SaveUser(app.User{token})
						app.StarOnMe(user, StarGoUrl)
					} else if c.Args().Len() < 1 {
						fmt.Println("Please enter you Github token. ", "Try again...")
					} else {
						fmt.Println("Just need only one Github token")
						fmt.Println("Please try again...")
					}
					return nil
				},
			},{
				Name:        "focus",
				Aliases:     []string{"f"},
				ArgsUsage: 	 "<owner/repository...>",
				Usage:       "Save the repository which you are interesting",
				Description: "The repository you want to focus on",
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 0 {
						repos := make([]app.RepoConfig, 0)
						for _, value := range c.Args().Slice() {
							repo , err := app.ParseToRepo(value)
							if err == nil {
								repos = append(repos, repo)
							}
						}
						app.SaveRepos(repos)

					} else {
						fmt.Println("Please try something like ''teyushen/star-go")
					}
					return nil
				},
			},{
				Name:        "append",
				Aliases:     []string{"a"},
				ArgsUsage: 	 "<owner/repository...>",
				Usage:       "Append the repository you are interesting on already focus",
				Description: "The repository you want to focus on",
				Action: func(c *cli.Context) error {
					if c.Args().Len() > 0 {
						repos := make([]app.RepoConfig, 0)
						for _, value := range c.Args().Slice() {
							repo , err := app.ParseToRepo(value)
							if err == nil {
								repos = append(repos, repo)
							}
						}
						app.AppendRepos(repos...)

					} else {
						fmt.Println("Please try something like ''teyushen/star-go")
					}
					return nil
				},
			},{
				Name:        "compare",
				Aliases:     []string{"c"},
				Usage:       "Use this to order the numbers of star",
				Description: "Compare all the repositories you are interesting",
				Action: func(c *cli.Context) error {
					app.CompareStar(app.PrepareReposInfo())
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}