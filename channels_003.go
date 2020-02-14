package main

import(
	"fmt"
)

func main() {
	
	dataStream := make(chan string)
	go func() {
		dataStream <- "Hi Peoples"
	}()
	
	/* Second Return on channel read 
	   true if read from write somewhere else in the process OR
	   true if reading a default value from closed channel
	*/
	
	data, ok := <- dataStream
	fmt.Printf("%v : %v \n", ok, data)
	
	
	/* Continued from 003.... */
	dataStream2 := make(chan int)
	go func() {
		//dataStream2 <- "Hi Peoples"
	}()	
	close(dataStream2)
	
	data2, ok2 := <- dataStream2
	fmt.Printf("%v : %d \n", ok2, data2)	
}