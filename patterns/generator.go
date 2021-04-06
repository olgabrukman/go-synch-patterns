package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

//WriterService is method that returns channel value.
//De-facto WriterService returns a channel that lets us communicate with the WriterService service it provides.

func RunSimple() {
	joe := WriterService("Joe")
	ann := WriterService("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func WriterService(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

