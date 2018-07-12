package app

import (
	"fmt"
)

func StarOnMe(u User, url string) {

	resp := request("PUT", u.Token, url)

	if resp.StatusCode == 204 {
		fmt.Println("Thanks for your star on me!")
	}
}

func RemoveStarOnMe(u User, url string) {

	resp := request("DELETE", u.Token, url)

	if resp.StatusCode == 204 {
		fmt.Println("Oh No! Why you do this to me?")
	}
}
