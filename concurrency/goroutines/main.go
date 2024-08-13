package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("first")
	go f("second")
	fmt.Println("done")
	time.Sleep(time.Second / 3)
}
