package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.github.com/users/ITV/repos"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var repos GithubRepos
	err = json.Unmarshal(responseData, &repos)
	if err != nil {
		fmt.Printf("%s", err)
	}
	for i, v := range repos {
		fmt.Println(i, v.Name)
	}
}
