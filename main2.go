package main

import (
	"sync"
	"time"
	"net/http"
)


func longProcess(wg *sync.WaitGroup, w http.ResponseWriter) {
	time.Sleep(5 * time.Second) //simulate work....
	
	w.Write([]byte("I am long process that finished... \n"))
	wg.Done();
}


func main() {
	var wg sync.WaitGroup
	
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		go longProcess(&wg, writer)
		
		wg.Add(1)
		wg.Wait()
		writer.Write([]byte("All Long Processes are gone \n"))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}	
}