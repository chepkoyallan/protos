package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)

	// reciever go routine/channel
	go func(ch <-chan int, wg *sync.WaitGroup){
		// if message was sent or not
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)
	
	// Sender go routine/ send only
	go func(ch chan<- int, wg *sync.WaitGroup){
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}