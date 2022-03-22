package main

import (
	"fmt"
	"time"
)

func f(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		// time.Sleep(3 * time.Second)
	}
}

func calculate(x, y int) {
	sum := x + y
	fmt.Println(sum)
	time.Sleep(3 * time.Second)
}

func main() {
	// go f("gorutine使って実行")
	// go calculate(3, 4)
	// f("普通に実行")
	// time.Sleep(3 * time.Second)
	// fmt.Println("done")

	messages := make(chan string)
	go func() { messages <- "Hello" }()

	msg := <-messages
	fmt.Println(msg)

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
