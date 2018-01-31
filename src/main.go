package main

import (
	"Repo-watcher/src/pkg/bodyMaker"
	"Repo-watcher/src/pkg/config"
	"Repo-watcher/src/pkg/github"
	"Repo-watcher/src/pkg/mailer"
	"Repo-watcher/src/pkg/timer"
	"fmt"
	"log"
	"time"
)

// Seperate Timer, Mailer, Watcher, BodyMaker
func appInit(c *config.ConfigFile) {
	// TODO: Watcher
	threeDaysAgo := time.Now().Add(-72 * time.Hour).Format("2006-01-02")
	option := github.NewListCommitsOptions(threeDaysAgo)
	commits, err := github.ListCommits(c, option)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	bm := bodyMaker.New(c)
	bm.AddCommitsToMessage(commits)
	body := bm.Body()

	m := mailer.New(c)
	if err := m.Send(body); err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("worked")
	// TODO: Timer
	for {
		c := config.FetchFromJson()
		if timer.IsTime(c, timer.GetCurrentJapanHourMin()) {
			appInit(c)
		}
		time.Sleep(1 * time.Minute)
	}
}
