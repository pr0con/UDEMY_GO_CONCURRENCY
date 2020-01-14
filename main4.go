package main

import(
	"fmt"
	"time"
	"sync"
)

type Object struct {
	Action *sync.Cond
}

//cd like in court is a conditional discharge... :)
func attachListener(cd *sync.Cond, fn func()) { 
	cd.L.Lock()
	defer cd.L.Unlock()
	
	cd.Wait()// wait for conditional discharge 
	fn() //run function provided
	
	//re-attaches event listener after fire
	go attachListener(cd, fn)
}

func main() {
	obj := Object{ Action: sync.NewCond(&sync.Mutex{} )} 

    
	go attachListener(obj.Action, func() {
		fmt.Println("Now I feel like a Javascript thing: Fire One");
	})
	
	go attachListener(obj.Action, func() {
		fmt.Println("Now I feel like a Javascript thing: Fire Two");
	})
	
	go attachListener(obj.Action, func() {
		fmt.Println("Now I feel like a Javascript thing: Fire Three");
	})
		
	
	for range time.Tick(time.Second * 2) {
		//obj.Action.Broadcast() //in order, just an observation maybe not everytime...
		
		//works prints out of order
		obj.Action.Signal()
		obj.Action.Signal()
		obj.Action.Signal()
	}
}

