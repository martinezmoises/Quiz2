package main

import "fmt"

//Two types of channels: Buffered and Unbuffered

// Normal method that is writting to a channel.
func message(msg chan string) {
	//Writing to the channel
	msg <- "HI, testing"
	//Important to close since if you don't close and you forget to read the data in the main it will give deadlock
	close(msg)
}

func main() {

	//msg is the channel that the sender method will use to transmit/send the data
	msg := make(chan string)
	//msg in the main go routine will have what sender function put it in
	go message(msg)
	//Now we have to read the data from the channel
	read := <-msg
	//Once the main recieves that data from the channel, then the function unblocks (continue the rest of the code)
	//printing the message
	fmt.Println(read)

	//Examples of buffered channels
	//This is a buffer channel, which means that go will put the values in that buffer (array)
	buffered := make(chan string, 4)
	//when creating a buffered channel, you wont need to block once they're in the capacity
	buffered <- "Hello"
	buffered <- "Hola"
	buffered <- "Todo bien pa?"
	buffered <- "SI"

	for i := 0; i < 4; i++ {
		fmt.Println(<-buffered)
	}

	//Example of an unbuffer (not saved)

	value := make(chan int)

	//Writing to the channel, blocking
	//value <- 22

	go func() {
		value <- 22
	}()

	//	This line won't get executed
	fmt.Println(<-value)

}
