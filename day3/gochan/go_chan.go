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
}
