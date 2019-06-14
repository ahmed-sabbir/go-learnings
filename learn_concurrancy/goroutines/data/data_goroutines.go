package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("Bob")
	wg.Wait()

	fmt.Println("Another one----")

	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg2.Done()
		}(i)
	}
	wg2.Wait()
}

func runMe(name string) {
	fmt.Println("hello", name, "from goroutine")
}
