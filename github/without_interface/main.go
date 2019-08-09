package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ReleasesInfo struct {
	Id      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

// Function to actually query thr Github API for the release information

func getLatestReleaseTag(repo string) (string, error) {
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

// Function to get the message to display ro the end user
func getReleaseTagMessage(repo string) (string, error) {
	tag, err := getLatestReleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error quering Github API: %s", err)
	}
	return fmt.Sprintf("The latest release is %s", tag), nil
}

func main() {
	msg, err := getReleaseTagMessage("laravel/framework")
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
	}
	fmt.Println(msg)
}
