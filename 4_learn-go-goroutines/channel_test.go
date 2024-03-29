package learn_go_goroutines

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // buffer 3
	defer close(channel)

	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "Hi"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Finish")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Loop " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive data", data)
	}

	fmt.Println("Finish")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		default:
			fmt.Println("Waiting for data")
		}
		if counter == 2 {
			break
		}
	}
}
