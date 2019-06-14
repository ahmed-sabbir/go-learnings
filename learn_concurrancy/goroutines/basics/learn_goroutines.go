package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		runMe()
		wg.Done()
	}()
	wg.Wait()
}

func runMe() {
	fmt.Println("Hello from a goroutine")
}
