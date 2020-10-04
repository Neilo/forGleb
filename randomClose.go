package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello() {
	for {
		time.Sleep(time.Second)
		fmt.Println("Привет")
	}
}

func main() {
	channel := make(chan string)
	rand.Seed(time.Now().UTC().UnixNano())
	go hello()
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(channel)
	}()

	fmt.Println(<-channel)
}
