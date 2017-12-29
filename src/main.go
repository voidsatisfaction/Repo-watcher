package main

import (
	"Repo-watcher/src/pkg/github"
	"Repo-watcher/src/pkg/mail"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func appInit() {
	c := getConfig()
	owners := "voidsatisfaction"
	repos := "TIL"
	twoDaysAgo := time.Now().Add(-48 * time.Hour).Format("2006-01-02")
	option := github.NewListCommitsOptions(twoDaysAgo)
	commits, err := github.ListCommits(owners, repos, option)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	msg := ""
	for i, commit := range *commits {
		if i == 0 {
			msg = fmt.Sprintf("URL %d: %s\n", i+1, commit.HTMLURL)
			continue
		}
		msg = fmt.Sprintf("%sURL %d: %s\n", msg, i+1, commit.HTMLURL)
	}

	fmt.Println(msg)
	m := &mail.Mail{
		From:     c.Mail.From,
		Username: c.Mail.Username,
		Password: c.Mail.Password,
		To:       c.Mail.To[0],
		Sub:      "TIL 복습 시간 입니다.",
		Msg:      msg,
	}

	if err := mail.GmailSend(m); err != nil {
		log.Println(err)
	}
}

func getCurrentJapanTime() string {
	// Time setting
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	nowJST := nowUTC.In(jst)
	hour, min, _ := nowJST.Clock()
	hourMin := ""
	if min < 10 {
		hourMin = fmt.Sprintf("%d:0%d", hour, min)
	} else {
		hourMin = fmt.Sprintf("%d:%d", hour, min)
	}
	return hourMin
}

type Config struct {
	Mail struct {
		From     string   `json:"from"`
		To       []string `json:"to"`
		Username string   `json:"username"`
		Password string   `json:"password"`
	} `json:"mail"`
	Github struct {
		Owner      string `json:"owner"`
		Repository string `json:"repository"`
	} `json:"github"`
	AlarmTime []string `json:"alarmTime"`
}

func getConfig() *Config {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	c := &Config{}
	json.Unmarshal(file, c)
	return c
}

func isTime(c *Config, t string) bool {
	for _, time := range c.AlarmTime {
		if time == t {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("worked")
	for {
		c := getConfig()
		if isTime(c, getCurrentJapanTime()) {
			appInit()
		}
		time.Sleep(1 * time.Minute)
	}
}
