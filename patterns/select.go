package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

func FanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:  c <- s
			case s := <-input2:  c <- s
			}
		}
	}()
	return c
}

func RunTimeoutWithSelect() {
	c := WriterService("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}
	}
}

func RunOneTimeoutWithSelect() {
	c := WriterService("Joe")
	//timeout created outside the loop
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}

func WriterServiceWithQuit(str string, quit chan string) chan string{ // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			select {
			case  c<- fmt.Sprintf("%s: %d", str, i):
				// do nothing
			case <-quit:
				//cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c // Return the channel to the caller.
}


func RunQuitWithSelect() {
	quit := make(chan string)
	c := WriterServiceWithQuit("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
	quit <- "Buy"
	fmt.Printf("Joe says: %q\n", <-quit)
}