package github

import (
	"Repo-watcher/src/pkg/config"
	"fmt"
	"net/url"
)

func createGithubRepositoryCommitsApiURL(owner, repos, since string) (*url.URL, error) {
	c := config.NewBasic()
	url, err := url.Parse(c.GithubApiHost)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("/repos/%s/%s/commits", owner, repos)
	url.Path = path
	q := url.Query()
	if since != "" {
		q.Add("since", since)
	}
	url.RawQuery = q.Encode()
	return url, nil
}
