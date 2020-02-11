package main
 
import(
	"fmt"
	"time"
	"context"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(context.TODO())
	
	time.AfterFunc(time.Second * 5, cancel)
	done := ctx.Done()
	
	for i := 0; ; i++ {
		select {
			case <-done:
				fmt.Println("Context Cancelled Program Ending...")
				return
			case <-time.After(time.Second):
				fmt.Println("tick: ",i)
		}
	}
}