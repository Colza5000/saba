package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://api.github.com/orgs/ITV/repos?access_token=< ACCESS TOKEN >&page=3" // All repos listed on the page specified
	// url := "https://api.github.com/repos/ITV/puppet-profile-traffic_director/releases?< ACCESS TOKEN >" // Repo specific info
	// url := "https://api.github.com/users/ITV/repos"
	// https://api.github.com/repos/ITV/puppet-consul/releases
	// url := "https://github.com/ITV/puppet-profile-node_service/releases"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	fmt.Println(responseString)
}
