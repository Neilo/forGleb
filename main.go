package main

import (
	"fmt"
)

func sort(x []int) []int {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if x[i] < x[j] {
				temp := x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	return x
}

func iter(x *int) {
	*x += 1
}

func main() {
	x := []int{12, 41, 21, 2, -4, 72, -1, 10}
	y := 1
	fmt.Println(x)
	fmt.Println(sort(x))
	fmt.Println(y)
	iter(&y)
	fmt.Println(y)
}
