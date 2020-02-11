package main

import (
	"os"
	"fmt"
	"time"
	"bytes"
	//"syscall"
	//"os/signal"
)

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		
		l, _ := os.Stdin.Read(data)
		
		data = bytes.Trim(data, "\x00")
		
		s := "\n"
		data = append(data, s...)
		
		len := int64(len(data))
		fmt.Println(len)
		
		if len > 2 && l > 1 {
			out <- data
		}
	}	
}

func main() {
	done := time.After(10 * time.Second)
	echo := make(chan []byte)
	
	go readStdin(echo)
	
	for {
		select {
			case buf := <-echo:
				os.Stdout.Write(buf)
			case <-done:
				fmt.Println("Timed out")
				close(echo)
				os.Exit(0)
			}
	}		
}