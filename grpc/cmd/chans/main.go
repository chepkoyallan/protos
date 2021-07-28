package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	// reciever go routine/channel
	go func(ch <-chan int, wg *sync.WaitGroup){
		// if message was sent or
		// if msg, ok := <-ch; ok {
		// 	fmt.Println(msg, ok)
		// }
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)
	
	// Sender go routine/ send only
	go func(ch chan<- int, wg *sync.WaitGroup){
		for i :=0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}