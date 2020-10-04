package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(ch chan string) {
	for {
		time.Sleep(time.Second)
		fmt.Println("Привет")
	}
}

func main() {
	//var wg sync.WaitGroup
	channel := make(chan string)
	rand.Seed(time.Now().UTC().UnixNano())
	go hello(channel)
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(channel)
	}()

	fmt.Println(<-channel)
}
