package chanPra

import "fmt"

func ChanTest() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)
}
