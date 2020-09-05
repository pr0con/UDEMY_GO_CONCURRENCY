package main

import(
	"os"
	"fmt"
	"time"
	
	"syscall"
	"os/signal"
)

/* Basically watch out for the nil value on a closed channel... */


func main() {
	data := make(chan bool)
	die := time.After(time.Second * 3)	


	go func(b chan bool) {
		for {
			select {
				case d := <- data: //Watch out for this... cause value will be nil on closed channel
					fmt.Println(d," Got Data ")
				case <- die:
					fmt.Println("Killing Go Routine...")
					return
				default:
					fmt.Println("chillin...")
					time.Sleep(time.Millisecond * 500)
			}
		}
	}(data)
	
		
	go func(b chan bool) {
		time.Sleep(time.Millisecond * 2)	
		data <- true
		close(data)		
	}(data)	
	
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGTSTP, syscall.SIGQUIT )
	
	for {
		<-c
		fmt.Println("killing program")
		break;
	}
}