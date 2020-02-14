package main

import(
	"fmt"
)

func main() {
	dataStream := make(chan int)
	
	go func() {
		defer close(dataStream)
		for i := 1; i <= 5; i++ {
			dataStream <- i
		}
	}()
	
	for i := range dataStream {
		fmt.Printf("%v \n", i)
	}
}