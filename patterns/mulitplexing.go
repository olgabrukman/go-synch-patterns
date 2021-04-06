package patterns

import "fmt"

/*fan in:
  actor 1 -----\
				multiplexor ---->
  actor 2 -----/
 */

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}
func RunMultiplexor() {
	c := FanIn(WriterService("Joe"), WriterService("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}