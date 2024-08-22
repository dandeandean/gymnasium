package main

import (
	"fmt"
	"sync"
)

func mystery(wg *sync.WaitGroup, s string, ch <-chan int) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("closed")
			return
		}
		fmt.Printf("%d", val)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go mystery(&wg, "Hello", ch)
	go mystery(&wg, "World", ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()

}
