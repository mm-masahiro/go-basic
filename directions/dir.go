package directions

func ReceivePing(pings chan<- string, msg string) {
	pings <- msg
}

func SendPing(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
