package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan string)
	
	go func() {
		for {
			select {
				case d := <-data:
					fmt.Println(d)
				case <-until:
					close(data)
					//time.Sleep(500 * time.Millisecond)
					return	
			}
		}
	}()
	
	for {
		data <- "lo"
		time.Sleep(500 * time.Millisecond)
	}
}