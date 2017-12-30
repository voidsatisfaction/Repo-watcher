package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

func GetConfig() *Config {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	c := &Config{}
	json.Unmarshal(file, c)
	return c
}
