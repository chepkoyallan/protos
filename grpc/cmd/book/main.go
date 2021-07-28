package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/chepkoyallan/protos/pkg/book"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main(){
	for i:=0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("From Cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("From Database")
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found with id: '%v\n'", id)
		time.Sleep(150 * time.Millisecond)
	}
	
}

func queryCache(id int) (book.Book, bool){
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int)(book.Book, bool){
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return book.Book{}, false
}