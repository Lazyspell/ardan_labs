package gochan

import (
	"fmt"
	"time"
)

func GoChan() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// BUG: All goroutines use the "i" for the for loop
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch //recieve
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg
		}
		close(ch)
	}()
	for msg := range ch {
		fmt.Println("got: ", msg)
	}
	msg = <-ch // ch is closed
	fmt.Printf("closed: %#v\n", msg)
}

/*
For every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- Send "n" over a channel
*/
func sleepSort(values []int) []int {

}

/* Channel semantics
- send & receive will block until opposite operation (*)
- receive from a closed channel will return a zero value without blocking
- send to a closed channel will panic
- closing a closed channel will panic
- send/receive to a nil channel will block forever
*/
