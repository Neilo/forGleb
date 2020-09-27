package main

import (
	"encoding/json"
	"fmt"
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
	const config = `
	{
		"port":":8888",
			"secrets":[
				{
					"user":"bot",
					"pass":"lehagleb"
				}
			],
			"db_credentials":{
				"user":"main",
				"pass":"cldyvcgaj",
				"url": "",
				"port":"5431",
				"data_base":"peoplerate"
			}
	}
`
	toByte := []byte(config)
	var cfg conf
	err := json.Unmarshal(toByte, &cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
