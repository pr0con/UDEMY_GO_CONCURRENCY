package main

import(
	"fmt"
	"time"
	"sync"
)

type Object struct {
	Action *sync.Cond
}

func main() {
	obj := Object{ Action: sync.NewCond(&sync.Mutex{} )} 

	//allows for recursive call inside variable func
	var attachListener func(cd *sync.Cond, fn func())

	//cd like in court is a conditional discharge... :)
	attachListener = func(cd *sync.Cond, fn func()) { 
			var wg sync.WaitGroup
			wg.Add(1) //make sure go routine is running 
			
			
			//launch go routine alert wait group after..
			go func() {
				wg.Done() //ok its running alert wait
				cd.L.Lock()
				defer cd.L.Unlock()
				
				cd.Wait()// wait for conditional discharge 
				fn() //run function provided
				
									
				//re-attaches event listener after fire
				go attachListener(cd, fn)
			}()
			
			wg.Wait() //wait for interal go func to fire exit this func
	}
	
	attachListener(obj.Action, func() {
		fmt.Println("Now I feel like a Javascript thing: Fire One");
	})
	
	attachListener(obj.Action, func() {
		fmt.Println("Now I feel like a Javascript thing: Fire Two");
	})
	
	attachListener(obj.Action, func() {
		fmt.Println("Now I feel like a Javascript thing: Fire Three");
	})


	
	for range time.Tick(time.Second * 2) {
		obj.Action.Broadcast()
		//obj.Action.Signal() //This will Signal One Two Three each second in order...
	}		
}