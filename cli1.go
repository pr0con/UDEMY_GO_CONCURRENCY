package main

import(
	"io"
	"os"
	"fmt"
	"bufio"
)

//use io.Reader as a string type
//<-chan read only string channel as a return type

func read(r io.Reader) (<-chan string) {
	lines := make(chan string)
	go func() {
		defer close(lines)
		scan := bufio.NewScanner(r)
		for scan.Scan() {
			lines <- scan.Text()
		}
	}()
	
	return lines
}

func main() {
	mes := read(os.Stdin)
	
	for anu := range mes {
		fmt.Println("Msg out: ",anu)
		switch(anu) {
			case "hello":
				fmt.Println("world");
				break;
			default:
				break;
		}
	}
}