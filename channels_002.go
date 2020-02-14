package main 

import(
	"fmt"
)

func main() {
	/* Internal POC ONE: Declaring a chan */
	
	var dataChan chan interface{}
	dataChan = make(chan interface{})
	
	go func() {
		dataChan <- "test"
	}()
	
	fmt.Println(<-dataChan)
	close(dataChan)
		
	
	/* Internal POC TWO: read write only chans */
	/*
	writeChan := make(chan <- interface{})
	readChan  := make(<- chan interface{})

	go func() {
		writeChan <- "test write read"
	}()
	go func() {
		readChan <- struct{}{}
	}()
	
	fmt.Println(<-writeChan)
	*/
	
	/* Internal POC Three: Accidental Channel Deadlock */
	
	dataChan2 := make(chan string)
	go func() {
		return  // return for any reason before sending causes dead lock
		dataChan2 <- "not happening"
	}()
	
	fmt.Println(<-dataChan2)
	
}