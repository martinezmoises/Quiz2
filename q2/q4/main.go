package main

import (
	"fmt"
	"sync" //FOr the use of waitgroups
)

// These functions will accept a pointer to a waitgroup
// Note that it is important we use pointer since we don't want to do a copy of the waitgroup
func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() //The.Done() function basically just decreaments wg by 1
	//FOr loop to print numbers from 1 - 5
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() //Also decreaments it by 1
	//Function to print letters a - d
	for i := 'a'; i <= 'd'; i++ {
		fmt.Printf("%c ", i)
	}
}
func main() {
	//Creating a waitgroup
	var wg sync.WaitGroup
	//We use the variable name.add to specify how many go routines you will be waiting for
	wg.Add(2)
	go printNumbers(&wg) //Need to dereference the pointer
	go printLetters(&wg)
	wg.Wait() //Will just wait for wg to decreament and reach 0
}
