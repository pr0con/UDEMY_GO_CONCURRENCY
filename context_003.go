package main
 
import(
	"fmt"
	"time"
	"context"
)
/* 
   decorator, specifies a time deadline
   if previous earlier than specified its ignored
   if done channel still open when deadline is met it gets closed
*/

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	time.AfterFunc(time.Second * 10, cancel)
	
	done := ctx.Done()
	for i := 0; ; i++ {
		select {
			case <-done:
				fmt.Println("Context Cancelled Program Ending...", ctx.Err())
				return
			case <-time.After(time.Second):
				fmt.Println("tick: ",i)
		}
	}
}