package main

import (
	"fmt"
	"main/patterns"
	"math/rand"
)

func main(){
	quit := make(chan string)
	c := patterns.WriterServiceWithQuit("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
	quit <- "Buy"
	fmt.Printf("Joe says: %q\n", <-quit)
}