package testFactory

import (
	"Repo-watcher/src/pkg/config"
	"Repo-watcher/src/test/util"
)

func NewConfig() *config.Config {
	c := &config.Config{}
	c.Mail.From = testUtil.MakeRandomString(20)
	c.Mail.To = []string{testUtil.MakeRandomString(20), testUtil.MakeRandomString(20)}
	c.Mail.Username = testUtil.MakeRandomString(10)
	c.Mail.Password = testUtil.MakeRandomString(10)

	c.Github.Owner = testUtil.MakeRandomString(15)
	c.Github.Repository = testUtil.MakeRandomString(15)

	c.AlarmTime = []string{testUtil.MakeRandomTimeString(), testUtil.MakeRandomTimeString()}
	return c
}
