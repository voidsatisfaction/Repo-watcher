package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Params struct {
	owners string
	repos  string
}

func NewParams(owners, repos string) *Params {
	return &Params{
		owners: owners,
		repos:  repos,
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

func ListCommits(owners, repos string, opt *ListCommitsOptions) (*RepositoryCommits, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owners, repos)

	if opt.since != "" {
		url = url + "?since=" + opt.since
	}
	fmt.Println(url)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
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
