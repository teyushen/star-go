package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	url := "https://api.github.com/users/teyushen/starred"
	basicAuth(url)
	//fmt.Printf("%s", basicAuth(url))
}

func basicAuth(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token 50c93442b8b519adfda8586a8f6a9418f4fac967")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	//fmt.Printf("%v", s)

	j := `[{"permissions":{
		      		"admin": true,
		      		"push": "true",
			        "pull": "true"
		    	}
		    }]`
	//x := make(map[string]interface{})

	//fmt.Println(string(bodyText))
	//json.Unmarshal([]byte(j), &x)
	//urlsJson, _ := json.Marshal(s)
	//fmt.Print(string(urlsJson))

	arr := make([]Station,0)
	fmt.Print(len(arr), cap(arr))
	repo := Station{}
	json.Unmarshal(bodyText, &repo)
	json.Unmarshal([]byte(j), &arr)

	fmt.Printf("%v\n", arr[0].Permissions.Admin)
	//fmt.Println(repo)

	/*for key, value := range x {
		fmt.Printf("%s -> %s\n", key, value)
	}*/

	return s
}

type repo struct {
	Admin bool `json:"admin"`
}
type Station struct {
	Permissions repo `json:"permissions"`
}