package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ReleaseInfoer interface {
	GetLatestReleaseTag(string) (string, error)
}

type ReleasesInfo struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
}
type GithubReleaseInfoer struct{}

// Function to actually query the GitHub API for the release information
func (gh GithubReleaseInfoer) GetLatestReleaseTag(repo string) (string, error) {

	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)

	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	releases := []ReleasesInfo{}

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	tag := releases[0].TagName

	return tag, nil

}

func getReleaseTagMessage(ri ReleaseInfoer, repo string) (string, error) {
	tag, err := ri.GetLatestReleaseTag(repo)

	if err != nil {
		return "", fmt.Errorf("Error quering GitHub API: %s", err)
	}
	return fmt.Sprintf("The latest release is %s", tag), nil
}

func main() {
	gh := GithubReleaseInfoer{}
	msg, err := getReleaseTagMessage(gh, "laravel/framework")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(msg)
}
