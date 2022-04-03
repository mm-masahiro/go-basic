package main

import (
	"fmt"
	"time"
)

type Number struct {
	num int
}

func (n Number) Add(i int) int {
	return n.num + i
}

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

	n := Number{num: 5}
	result := n.Add(5)

	const hoge = 10 / 2
	foo := 20

	fmt.Println(n.Add(hoge))
	fmt.Println(n.Add(foo))
	fmt.Println(result)

	// messages := make(chan string)
	// go func() { messages <- "Hello" }()

	// msg := <-messages
	// fmt.Println(msg)

	// 第二引数でサイズを指定できる
	// ch := make(chan int, 3)
	// ch <- 1
	// ch <- 2
	// ch <- 3

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// chanPra.ChanTest()
	// chanPra.ChanTest2()

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)

	// directions.ReceivePing(pings, "Hello")
	// directions.SendPing(pings, pongs)
	// fmt.Println(<-pongs)
}
