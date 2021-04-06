package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Str string
	Wait chan bool
}

func WriterServiceMsg(str string) <-chan Message { // Returns receive-only channel of strings.
	c := make(chan Message)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- Message{
				Str: fmt.Sprintf("%s %d", str, i), Wait: make(chan bool),
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func FanInMsgWithSequenceOrder(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}

func RunRestoreMsgSequencing() {
	c := FanInMsgWithSequenceOrder(WriterServiceMsg("Joe"), WriterServiceMsg("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
	for i := 0; i < 5; i++ {
		msg1 := <-c;
		fmt.Println()
		msg2 := <-c;
		fmt.Println(msg2.Str)
		msg1.Wait <- true
		msg2.Wait <- true
	}
}


