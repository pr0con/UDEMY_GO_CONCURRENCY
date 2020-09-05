package main 

import(
	"fmt"
	"time"
)


func main() {
	/* READ NIL BLOCK / DEADLOCK, Example 1 */
	/*
		var dataChan chan interface{}
		<-dataChan
	*/
	
	/* CLOSE NIL PANIC */
	/*
		var dataChan chan interface{}
		close(dataChan)		
	*/


	/* some characteristics: 
		- return read only chan
		- encapsulated go func	
	*/
	funcVar := func() <-chan int {
		irs := make(chan int, 5) //instantiated, written to, closed from here...
		go func() {
			defer close(irs)
			for i := 0; i <= 5; i++ {
				irs <- i
			}
		}()
		return irs
	}

	exrs := funcVar()
	for res := range exrs {
		fmt.Printf("Recieved: %d\n",res)
	}
	
	/* 
		Unlike A Switch statement using select actually listens to all case at the same time 
		The whole select statement will block if nothing is ready
	*/	
	chanOne := make(chan interface{}); close(chanOne)
	chanTwo := make(chan interface{}); close(chanTwo)
	
	var countOne, countTwo int
	for i := 1; i <= 1000; i++ {
		select {
			case <-chanOne:
				countOne++
			case <-chanTwo:
				countTwo++
		}
	}
	fmt.Printf("Count One: %d \nCount Two: %d \n", countOne, countTwo)
	
	
	
	/* Blocking Select Example :: never unblocks reading from nil channel*/	
	var chanThree <-chan int
	select {
		case <-chanThree:
		case <-time.After(time.Second * 3):
			fmt.Println("read nil timed out")
	}
	
	
	fmt.Println("Program Done. But not dead...")	
	//select{}	
}
