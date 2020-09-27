package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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
