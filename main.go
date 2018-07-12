package main

import (
	"log"
	"io/ioutil"
	"github.com/teyushen/star-go/app"
)

const StarGoUrl = "https://api.github.com/user/starred/teyushen/star-go"
const StarredUrl = "https://api.github.com/users/teyushen/starred"

func init() {
	log.SetOutput(ioutil.Discard)
}


func main() {

	//user := app.GetUser()
	//app.RemoveStarOnMe(user, StarGoUrl)


	//url := "https://api.github.com/users/teyushen/starred"
	//ur1 := "https://api.github.com/users/believeWang/starred"
	//url1 := "https://api.github.com/user/starred/teyushen/star-go"


	app.SaveRepo(app.Repo{})

	//fmt.Print(app.GetRepoInfo(user, url))
	//fmt.Print(app.CollectAllReposInfo(user, url, ur1))
	//app.CompareStar(app.CollectAllReposInfo(user, url, ur1))


	//app := &cli.App{
	//	Name:        "star-go",
	//	Version:     "0.1.0",
	//	Description: "This can help us easily compare how many stars of Repository in Github we interested",
	//	Authors: []*cli.Author{
	//		{Name: "Dennis", Email: "teyushen@gmail.com"},
	//	},
	//	Flags: []cli.Flag{
	//		&cli.StringFlag{Name: "name", Value: "bob", Usage: "a name to say"},
	//	},
	//	Commands: []*cli.Command{
	//		{
	//			Name:        "init",
	//			Aliases:     []string{},
	//			UsageText: 	 "star-go init TOKEN",
	//			Usage:       "use it to initial the basic configuration",
	//			Description: "This is how we add Github token into our config",
	//			Action: func(c *cli.Context) error {
	//				if c.Args().Len() == 1 {
	//					token := c.Args().First()
	//					fmt.Println("Your Github token is: ", c.Args().First())
	//					app.SaveUser(app.User{token})
	//					app.StarOnMe(user, StarGoUrl)
	//				} else {
	//					fmt.Println("Just need only one Github token")
	//					fmt.Println("Please try again...")
	//				}
	//				return nil
	//			},
	//		},
	//	},
	//}
	//app.Run(os.Args)

	//u := app.User{}
	//u.Token = ""

	//app.RemoveStarOnMe(u, url1)
	//app.CompareStar(user, url)

}