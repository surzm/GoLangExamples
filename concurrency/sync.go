package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("1st goroutine sleeping...")
		time.Sleep(100 * time.Millisecond)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(200 * time.Millisecond)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")

	fmt.Println("Next example")
	go func() {
		wg.Add(1) //
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()
	wg.Wait()
	fmt.Println("All goroutines complete.")
}
