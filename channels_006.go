package main

/* Cleaning up goroutines and channels
		- if not cleaned up this can cause leaks
*/

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan string)
	until := time.After(5 * time.Second)
	done := make(chan bool) // apply fix
	
	go func() {
		for {
			select {
				case d := <-data:
					fmt.Println(d)
				case <-until:
					done <- true // apply fix
					//close(data)  // origonal
					return
			}
		}
	}()
	
	/*
	for {
		data <- "lo"
		time.Sleep(500 * time.Millisecond)
	}
	*/
	
	/* Fix */	
	for {
		select {
			case <-done:
				fmt.Println("Done")
				close(data)
				return
			default:
				data <- "lo"
				time.Sleep(500 * time.Millisecond)
			
		}
	}
	
}