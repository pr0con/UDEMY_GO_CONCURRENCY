package main

import (
	"fmt"
	"sync"
	"time"
	"net/http"
)


func longProcess(wg *sync.WaitGroup, w http.ResponseWriter) {
	//time.Sleep(5 * time.Second) //simulate work....
    
    
    for {
        fmt.Println("Infinite Loop 1")
        time.Sleep(time.Second)
    }	
	
	
	/*
    for true {
        fmt.Println("Infinite Loop 2")
        time.Sleep(time.Second)
    }		
	*/
	
	w.Write([]byte("I am long process that finished... \n"))
	wg.Done();
}


func main() {
	var wg sync.WaitGroup
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go longProcess(&wg, w)
		
		wg.Add(1)
		wg.Wait()
		w.Write([]byte("All Long Processes are gone \n"))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}	
}