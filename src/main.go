package main

import (
	"Repo-watcher/src/pkg/config"
	"Repo-watcher/src/pkg/github"
	"Repo-watcher/src/pkg/mailer"
	"Repo-watcher/src/pkg/timer"
	"fmt"
	"log"
	"time"
)

// Seperate Timer, Mailer, Watcher, TemplateEngine
func appInit(c *config.Config) {
	owner := c.Github.Owner
	repos := c.Github.Repository
	// TODO: Watcher
	threeDaysAgo := time.Now().Add(-72 * time.Hour).Format("2006-01-02")
	option := github.NewListCommitsOptions(threeDaysAgo)
	commits, err := github.ListCommits(owner, repos, option)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// TODO: Mailer
	msg := ""
	for i, commit := range *commits {
		if i == 0 {
			msg = fmt.Sprintf("URL %d: %s\n", i+1, commit.HTMLURL)
			continue
		}
		msg = fmt.Sprintf("%sURL %d: %s\n", msg, i+1, commit.HTMLURL)
	}

	fmt.Println(msg)
	m := &mailer.Mail{
		From:     c.Mail.From,
		Username: c.Mail.Username,
		Password: c.Mail.Password,
		To:       c.Mail.To[0],
		Sub:      "TIL 복습 시간 입니다.",
		Msg:      msg,
	}

	if err := mailer.GmailSend(m); err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("worked")
	// TODO: Timer
	for {
		c := config.GetConfig()
		if timer.IsTime(c, timer.GetCurrentJapanHourMin()) {
			appInit(c)
		}
		time.Sleep(1 * time.Minute)
	}
}
