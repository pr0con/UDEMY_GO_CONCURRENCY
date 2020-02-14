package main

/*
	type Context inteface {
		Deadline() (deadline time.Time, ok bool)	
		Done() <-chan struct{}
		Err() error
		Value(key interface{}) interface{}
	}	
*/

import(
	"fmt"
	"time"
	"context"
)


func main() {
	ctx := context.Background()
	done := ctx.Done()
	
	for i := 0; ;i++ {
		select {
			case <-done:
				return
			case <-time.After(time.Second):
				fmt.Println("tick: ",i)
		}
	}
	fmt.Println("never gets here")
}
