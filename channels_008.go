package main 

import(
	"os"
	"fmt"
	"time"
	"syscall"
	"os/signal"	
)

/* 
	Buffered Channel unblocking until full given there is space for one we can block till work is done...
*/	

func main() {
	lock := make(chan bool, 1)
	
	for i := 1; i <= 5; i++ {
		go func(i int) {
			for {
				lock <- true
				fmt.Printf("func %d working...\n", i)
				time.Sleep(time.Second * 2)
				fmt.Printf("func %d done working...\n", i)
				<-lock
			}
		}(i)
	}
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGTSTP, syscall.SIGQUIT )
	
	for {
		<-c
		fmt.Println("killing program")
		break;
	}	
}	
