package main

import (
	"fmt"
	"sync"
)

// Made two goroutines that send numbers to channels - one sends only even numbers, other odd numbers.
// TASK: Read them all and print to stdout.
func main() {
	messages := make(chan int)

	go func() {
		defer close(messages)

		var wg sync.WaitGroup

		wg.Add(2)

		go func() {
			messages <- 1
			messages <- 3
			messages <- 5
			messages <- 7
			messages <- 9

			wg.Done()
		}()

		go func() {
			messages <- 2
			messages <- 4
			messages <- 6
			messages <- 8
			messages <- 10

			wg.Done()
		}()

		wg.Wait()
	}()

	for v := range messages {
		fmt.Println(v)
	}
}
