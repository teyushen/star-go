package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"encoding/json"
)

func main() {
	url := "https://api.github.com/repos/teyushen/oauth2/stargazers"
	basicAuth(url)
	//fmt.Printf("%s", basicAuth(url))
}

func basicAuth(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token e229b32f20ae12c500edfee36dd2d69b81b49b6b")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	//fmt.Printf("%v", s)

	j := `[{"permissions":{
		      		"admin": "true",
		      		"push": "true",
			        "pull": "true"
		    	}
		    }]`
	x := make(map[string]interface{})

	json.Unmarshal([]byte(j), &x)

	//urlsJson, _ := json.Marshal(j)
	//fmt.Print(string(urlsJson))

	//fmt.Printf("%v\n", string(j[0]))
	fmt.Printf("%v\n", x)
	text := `{"people": [{"craft": "ISS", "name": "Sergey Rizhikov"}, {"craft": "ISS", "name": "Andrey Borisenko"}, {"craft": "ISS", "name": "Shane Kimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"}, {"craft": "ISS", "name": "Thomas Pesquet"}, {"craft": "ISS", "name": "Peggy Whitson"}], "message": "success", "number": 6}`
	textBytes := []byte(text)

	people1 := people{}
	json.Unmarshal(textBytes, &people1)
	fmt.Println(people1.Message)
	/*for key, value := range x {
		fmt.Printf("%s -> %s\n", key, value)
	}*/

	return s
}

type people struct {
	Message string `json:"message"`
}
