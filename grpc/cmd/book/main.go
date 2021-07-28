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
	m := &sync.Mutex{}

	for i:=0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex){
			if b, ok := queryCache(id, m); ok {
				fmt.Println("From Cache")
				fmt.Println(b)
				
			}
			wg.Done()
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.Mutex){
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("From Database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		
		// fmt.Printf("Book not found with id: '%v\n'", id)
		// time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
	
}

func queryCache(id int, m *sync.Mutex) (book.Book, bool){
	m.Lock()
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}

func queryDatabase(id int, m *sync.Mutex)(book.Book, bool){
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return book.Book{}, false
}