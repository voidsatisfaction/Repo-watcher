package bodyMaker

import (
	"Repo-watcher/src/pkg/config"
	"Repo-watcher/src/pkg/github"
	"fmt"
	"strings"
)

type BodyMaker struct {
	To      []string
	Repos   []string
	Subject string
	Message string
}

func New(c *config.ConfigFile) *BodyMaker {
	bm := &BodyMaker{
		To:    c.Mail.To,
		Repos: []string{c.Github.Repository},
	}
	bm.Subject = strings.Join(bm.Repos, " ") + " 복습 시간 입니다."

	return bm
}

func (bm *BodyMaker) AddCommitsToMessage(rcs *github.RepositoryCommits) {
	var msg string
	for i, commit := range *rcs {
		if i == 0 {
			msg = fmt.Sprintf("URL %d: %s\n", i+1, commit.HTMLURL)
			continue
		}
		msg = fmt.Sprintf("%sURL %d: %s\n", msg, i+1, commit.HTMLURL)
	}
	bm.Message = msg
}

func (bm *BodyMaker) Body() string {
	return fmt.Sprintf(
		"%s%s%s",
		bm.createTo(),
		bm.createSubject(),
		bm.createBody(),
	)
}

func (bm *BodyMaker) createTo() string {
	return fmt.Sprintf("To: %s\r\n", strings.Join(bm.To, " "))
}

func (bm *BodyMaker) createSubject() string {
	return fmt.Sprintf("Subject: %s\r\n\r\n", bm.Subject)
}

func (bm *BodyMaker) createBody() string {
	message := bm.Message
	return message
}
