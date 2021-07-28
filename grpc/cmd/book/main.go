package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/chepkoyallan/protos/pkg/book"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main(){
	//Wait groups
	wg := &sync.WaitGroup{}

	for i:=0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup){
			if b, ok := queryCache(id); ok {
				fmt.Println("From Cache")
				fmt.Println(b)
				
			}
			wg.Done()
		}(id, wg)
		go func(id int, wg *sync.WaitGroup){
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From Database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
		
		// fmt.Printf("Book not found with id: '%v\n'", id)
		// time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
	
}

func queryCache(id int) (book.Book, bool){
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int)(book.Book, bool){
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			// cache[id] = b
			return b, true
		}
	}
	return book.Book{}, false
}