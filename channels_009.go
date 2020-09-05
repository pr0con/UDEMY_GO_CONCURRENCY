package main

import(
	"fmt"
)

func main() {
	alphas := []rune{'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}
	fmt.Println(string(alphas))
	
	
	ctrlVal := 10
	clearBuf := make(chan bool)
	
	bufChanData := make(chan rune, ctrlVal)
	go func(bcd chan rune) {
		for {
			select {
				case cb := <- clearBuf:
					for i := 0; i <= ctrlVal; i++ {
						fmt.Println(cb, <-bcd)
					}
			}
		}
	}(bufChanData)
	
	
	for i := 0; i <= ctrlVal; i++ {
		if i == ctrlVal {
			clearBuf <- true
			break;
		}
		bufChanData <- alphas[i]
		fmt.Println("Added Rune ", string(alphas[i]))
	}
	
}