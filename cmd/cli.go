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
					if c.NArg() == 1 {
						token := c.Args().First()
						fmt.Println("Your Github token is: ", c.Args().First())
						user := app.User{token}
						app.SaveUser(user)
						app.StarOnMe(user, StarGoUrl)
					} else if c.NArg() < 1 {
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
				ArgsUsage: 	 "<owner/repository>...",
				Usage:       "Save the repository which you are interested",
				Description: "The repository you want to focus on",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "tab", Aliases: []string{"t"}, Value: 1, Usage: "which tab you want to save"},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						repos := make([]app.RepoConfig, 0)
						for _, value := range c.Args().Slice() {
							repo , err := app.ParseToRepo(value)
							if err == nil {
								repos = append(repos, repo)
							}
						}
						app.SaveRepos(repos, c.Int("tab"))

					} else {
						fmt.Println("Please try something like ''teyushen/star-go")
					}
					return nil
				},
			},{
				Name:        "append",
				Aliases:     []string{"a"},
				ArgsUsage: 	 "<owner/repository>...",
				Usage:       "Append the repository you are interested on already focus",
				Description: "The repository you want to focus on",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "tab", Aliases: []string{"t"}, Value: 1, Usage: "which tab you want to save"},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() > 0 {
						repos := make([]app.RepoConfig, 0)
						for _, value := range c.Args().Slice() {
							repo , err := app.ParseToRepo(value)
							if err == nil {
								repos = append(repos, repo)
							}
						}
						app.AppendRepos(c.Int("tab"), repos...)

					} else {
						fmt.Println("Please try something like ''teyushen/star-go")
					}
					return nil
				},
			},{
				Name:        "compare",
				Aliases:     []string{"c"},
				Usage:       "Use this to order the numbers of star",
				Description: "Compare all the repositories you are interested",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "tab", Aliases: []string{"t"}, Value: 1, Usage: "which tab you want to save"},
				},
				Action: func(c *cli.Context) error {
					app.CompareStar(app.PrepareReposInfo(c.Int("tab")))
					return nil
				},
			},{
				Name:        "list",
				Aliases:     []string{"ls"},
				Usage:       "Display all the focus repositories",
				Description: "List all the repo you are interested",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "tab", Aliases: []string{"t"}, Value: 1, Usage: "which tab you want to save"},
				},
				Action: func(c *cli.Context) error {
					//cc := color.New(color.FgBlue)
					fmt.Printf("%10s %20s\n", "Owner", "Repositories")
					fmt.Println("-----------------------------------------------------------------------")
					for _, repo := range app.GetRepos(c.Int("tab")) {
						fmt.Printf("%10s        %s\n", repo.Owner, repo.RepoNames)
					}
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}