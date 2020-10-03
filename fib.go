package main

import (
	"fmt"
	"sync"
	"time"
)

func fib(num int, wg *sync.WaitGroup) int {
	wg.Add(1)
	if num == 0 || num == 1 {
		return 1
	}
	defer wg.Done()
	return fib(num-2, wg) + fib(num-1, wg)
}

func main() {
	var wg sync.WaitGroup
	go fmt.Println(fib(25, &wg))
	time.Sleep(time.Second)
	go fmt.Println(fib(20, &wg))
	time.Sleep(time.Second)
	wg.Wait()
}
