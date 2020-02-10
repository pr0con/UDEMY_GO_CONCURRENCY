package main

import(
	"fmt"
	"net"
	"sync"
	"bufio"
	"strings"	
)


func handleData(c net.Conn, l net.Listener) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil { fmt.Println(err); return }
		
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" { break; }
		
		c.Write([]byte("Hi there from server...\n"))	
	}
	c.Close()
	l.Close()		
}

func handleConnections(port string, wg *sync.WaitGroup) {
	l, err := net.Listen("tcp4", ":"+port); if err != nil { fmt.Println(err); return }
	
	for {
		c, err := l.Accept(); if err != nil { fmt.Println("Listener Killed...\n", err, "\n"); break; } else {
			go handleData(c, l)
		}
	}
	fmt.Println("Handle Connections Ended for ",port);
		
	wg.Done()	
}


func main() {
	var wg sync.WaitGroup
	
	wg.Add(1)
	go handleConnections("1337", &wg)
	wg.Add(1)
	go handleConnections("1338", &wg)
	wg.Add(1)
	go handleConnections("1339", &wg)
	wg.Wait()
	fmt.Println("Server done handling connections")
}