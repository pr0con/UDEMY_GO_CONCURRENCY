package main 

import(
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()
	
	cycles := 0
	start  := time.Now()	
		
	Loop:
		for {
			select {
				case <-done:
					break Loop
				default:
					fmt.Printf("Cycles: Unblocked for %v \n", time.Since(start))
			}
			
			/* Do some other work here */
			cycles++
			time.Sleep(time.Second * 1)
		}
	
	fmt.Printf("Did %d cycles of work. \n", cycles)	
}

