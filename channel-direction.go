package main

import "fmt"

func send(pings chan<- string, msg string) {
    pings <- msg
}

func receive(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    send(pings, "sent message")
    receive(pings, pongs)
    fmt.Println(<-pongs)
}
