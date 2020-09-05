package main 

/* Compare with channels_009.go */

import(
	"os"
	"fmt"
	"bytes"
)

func main() {	
	var OBuf bytes.Buffer
	defer OBuf.WriteTo(os.Stdout)
	
	alphas := []rune{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}
		
	ctrlVal := 25	
	bufChanData := make (chan rune, ctrlVal)
	
	go func(bcd chan rune, ra []rune, bb *bytes.Buffer) {
		defer close(bcd)	
		for i := 0; i <= ctrlVal; i++ {
			fmt.Fprintf(bb, "SENT RUNE: %v \n", string(ra[i]))
			bcd<-ra[i]
		}
	}(bufChanData, alphas, &OBuf)	
	
	
	for Rune := range bufChanData {			
		fmt.Fprintf(&OBuf, "GOT RUNE: %v \n", string(Rune))
	}
}