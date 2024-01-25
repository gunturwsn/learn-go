package learn_go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
	}()

	data := <-channel
	fmt.Println(data)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	/* channel as a parameter by default like pass by pointer  */
	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World!"
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)  // send data to the channel
	go OnlyOut(channel) // receive data from the channel

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(3 * time.Second)
	channel <- "Hello World!!"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
