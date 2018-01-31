package github

import (
	"Repo-watcher/src/pkg/config"
	"Repo-watcher/src/test/util"
	"fmt"
	"testing"
	"time"
)

func TestCreateGithubApiURL(t *testing.T) {
	tests := []struct {
		owner string
		repos string
		since string
	}{
		{
			testUtil.MakeRandomString(10),
			testUtil.MakeRandomString(10),
			time.Now().Add(-72 * time.Hour).Format("2006-01-02"),
		},
	}

	for _, test := range tests {
		owner, repos, since := test.owner, test.repos, test.since
		expect := fmt.Sprintf("%s/repos/%s/%s/commits?since=%s", config.NewBasic().GithubApiHost, owner, repos, since)
		url, err := createGithubRepositoryCommitsApiURL(owner, repos, since)
		if err != nil {
			t.Errorf("Error occured!")
			t.Errorf("error: %+v", err)
		}
		got := url.String()
		if got != expect {
			t.Errorf("Expect: %s, Got: %s", expect, got)
		}
	}
}
