package main

import (
	"fmt"
	"time"
)

func main() {

	chInt := make(chan int)

	//<-chInt // deadlock
	//chInt <-1 // deadlock

	go func() {
		for i := range chInt {
			fmt.Println("from chan: ", i)
		}
	}()
	for i := 0; i < 5; i++ {
		//i := i
		go func() {
			chInt <- i
		}()
	}

	time.Sleep(1 * time.Second)
	close(chInt)
	//chInt <- 1 // panic

	i := <-chInt
	fmt.Println(i) // 0 - default value

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
