package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func hello(ch chan string, wg *sync.WaitGroup) {
	for {
		time.Sleep(time.Second)
		fmt.Println("Привет")
	}
	ch <- "Хеллов"
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan string)
	rand.Seed(time.Now().UTC().UnixNano())
	go hello(channel, &wg)
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		wg.Wait()
		close(channel)
	}()

	fmt.Println(<-channel)
}
