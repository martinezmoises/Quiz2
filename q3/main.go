package main

import (
	"fmt"
	"time"
)

//GO routines are basically independent entity

func printMessage() {
	fmt.Println("Hi. This is displaying from 'printMessage' ")
}

func printMessage2() {
	fmt.Println("Hello. This is displaying from 'printMessage2'")
}

// Main is a go routine
// Have to keep in mind that main is our diver, therefore everything will start here at main
func main() {
	//Once we add the go and it becomes a go routine, we lose the idea of determinism because we cannot accuratley say which one will run first because they
	//are now independent functions
	go printMessage()
	go printMessage2()
	//Now when you run the code using go routines, main won't wait for the functions anymore because they are now independent
	//Therefore it exits.

	//To fix this we let main sleep(Delay the exiting of the program)
	time.Sleep(5 * time.Second)
}

//Using sleep is not recommended, thats why we use channels and such
