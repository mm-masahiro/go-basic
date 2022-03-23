package chanPra

import "fmt"

func ChanTest() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)
}

func ChanTest2() {
	numbers := make(chan int)

	go func() { numbers <- 3 }()

	num := <-numbers

	fmt.Println(num)
}
