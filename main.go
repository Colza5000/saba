package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type GithubRepos []struct {
	URL             string `json:"url"`
	AssetsURL       string `json:"assets_url"`
	UploadURL       string `json:"upload_url"`
	HTMLURL         string `json:"html_url"`
	ID              int    `json:"id"`
	NodeID          string `json:"node_id"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name            string `json:"name"`
	Draft           bool   `json:"draft"`
	Author          struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Prerelease  bool          `json:"prerelease"`
	CreatedAt   time.Time     `json:"created_at"`
	PublishedAt time.Time     `json:"published_at"`
	Assets      []interface{} `json:"assets"`
	TarballURL  string        `json:"tarball_url"`
	ZipballURL  string        `json:"zipball_url"`
	Body        string        `json:"body"`
}

func main() {
	// url := "https://api.github.com/orgs/ITV/repos?access_token=< ACCESS TOKEN >&page=3" // All repos listed on the page specified
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
	//
	// responseString := string(responseData)

	// fmt.Println(responseString)

	var repos GithubRepos
	err = json.Unmarshal(responseData, &repos)
	if err != nil {
		fmt.Printf("%s", err)
	}
	// fmt.Printf("%+v", repos)
	for _, v := range repos {
		// fmt.Printf("%s ", v.TagName)
		intSlice := []string{v.TagName}
		// first := intSlice[1:1]
		fmt.Println(intSlice)
	}
}
