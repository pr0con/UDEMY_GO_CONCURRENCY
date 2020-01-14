package main

import (
	//Native Go
	"fmt"
	"time"
	"sync"
	//"runtime"
	
	//3rd Party
	
	
	//Our Packages
)

func init() {
	fmt.Println("Initializing Go Application");
}

func longProcess(wg *sync.WaitGroup) {
    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
    time.Sleep(2 * time.Second) //simulate work....
    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
    
    fmt.Println("Work Finished")
    wg.Done()		
}


func main() {
	
	var wg sync.WaitGroup
	
	fmt.Println("Go Program Running");


	go func() {
		for range time.Tick(time.Second * 2) {
			fmt.Println("Engine #1 is working: ")
		}
	}()
	
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go longProcess(&wg)
	}
	
	
	wg.Wait()
	fmt.Println("Done Waiting");
	
	/*		
	for range time.Tick(time.Second * 2) {
		fmt.Println("Engine #1 is working: ", runtime.NumGoroutine(), " tasks(Go Routines) running");
	}
	*/
}







