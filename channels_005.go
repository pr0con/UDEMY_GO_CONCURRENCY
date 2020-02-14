package main

import(
	"fmt"
	"time"
)

func main() {
	start := make(chan interface{})
	
	go func() {
		<-start
		fmt.Println("go func 1")
	}()
	
	go func() {
		<-start
		fmt.Println("go func 2")
	}()
	
	go func() {
		<-start
		fmt.Println("go func 3")
	}()		
	
	<-time.After(time.Second * 5)
	close(start)
	<-time.After(time.Second * 5)
}