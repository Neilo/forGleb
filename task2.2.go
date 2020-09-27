package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type conf struct {
	Port          string   `json:"port"`
	Secrets       []secret `json:"secrets"`
	Dbcredentials dbcred   `json:"db_credentials"`
}

type dbcred struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	URL      string `json:"url"`
	Port     string `json:"port"`
	Database string `json:"data_base"`
}

type secret struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}
	var cfg conf
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
