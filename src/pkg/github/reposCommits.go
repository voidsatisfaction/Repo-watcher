package github

import (
	"Repo-watcher/src/pkg/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Params struct {
	owner string
	repos string
}

func NewParams(owner, repos string) *Params {
	return &Params{
		owner: owner,
		repos: repos,
	}
}

type ListCommitsOptions struct {
	since string
	until string
}

func NewListCommitsOptions(since string) *ListCommitsOptions {
	return &ListCommitsOptions{
		since: since,
	}
}

type RepositoryCommits []RepositoryCommit

type RepositoryCommit struct {
	SHA         string `json:"sha,omitempty"`
	HTMLURL     string `json:"html_url,omitempty"`
	URL         string `json:"url,omitempty"`
	CommentsURL string `json:"comments_url,omitempty"`
}

func ListCommits(c *config.ConfigFile, opt *ListCommitsOptions) (*RepositoryCommits, error) {
	owner, repos, since := c.Github.Owner, c.Github.Repository, opt.since
	url, err := createGithubRepositoryCommitsApiURL(owner, repos, since)
	if err != nil {
		return nil, err
	}
	fmt.Println(url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Repo-watcher-app")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rcs := &RepositoryCommits{}
	json.Unmarshal(body, rcs)

	return rcs, nil
}
