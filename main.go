package main

import "github.com/teyushen/star-go/app"

func main() {
	//url := "https://api.github.com/users/teyushen/starred"
	url1 := "https://api.github.com/user/starred/teyushen/star-go"
	//app.BasicAuth(url)
	//app.StarOnMe(url1)
	app.RemoveStarOnMe(url1)
	//fmt.Printf("%s", basicAuth(url))
}

