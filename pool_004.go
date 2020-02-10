package main

import (
	"os"
	"fmt"
	"net"
	"sync"
	
	"bufio"
	"strings"
	"syscall"
	"os/signal"
)


type ConnectionPool struct {
	mutex sync.RWMutex
	list  map[int]net.Conn
}


func NewConnectionPool() *ConnectionPool {
	pool := &ConnectionPool{
		list: make(map[int]net.Conn),
	}
	return pool
}

func (pool *ConnectionPool) Add(c net.Conn) int {
	pool.mutex.Lock()
	nextConnectionId := len(pool.list)
	pool.list[nextConnectionId] = c
	pool.mutex.Unlock()
	return nextConnectionId
}

func (pool *ConnectionPool) Get(connectionId int) net.Conn {
	pool.mutex.RLock()
	c := pool.list[connectionId]
	pool.mutex.RUnlock()
	return c
}


func (pool *ConnectionPool) Remove(connectionId int) {
	pool.mutex.Lock()
	delete(pool.list, connectionId)
	pool.mutex.Unlock()
}

func (pool *ConnectionPool) Size() int {
	return len(pool.list)
}

func (pool *ConnectionPool) Range(callback func(net.Conn, int)) {
	pool.mutex.RLock()
	for cid, c := range pool.list {
		callback(c, cid)
	}
	pool.mutex.RUnlock()
}

//Additions...
func handleData(nc net.Conn) {
	for {
		netData, err := bufio.NewReader(nc).ReadString('\n')
		if err != nil { fmt.Println(err); return }
		
		client_msg := strings.TrimSpace(string(netData))
		fmt.Println("Client Wrote: ", client_msg)
		
		if client_msg == "QUIT" { nc.Close();  break;}
	}
}

func main() {
	socket, err := net.Listen("tcp", "127.0.0.1:1337")
	if err != nil { fmt.Println(err) }
	
	connectionPool :=  NewConnectionPool()

	go func(pool *ConnectionPool) {
		for {
			c, _ := socket.Accept()
			
			cid := pool.Add(c)
			fmt.Println("New Client ID: ", cid)
			
			size := pool.Size()
			fmt.Println("Pool Size: ", size)
			
			go handleData(pool.Get(cid))
			
			pool.Range(func(targetConnection net.Conn, targetConnectionId int) {
				writer := bufio.NewWriter(targetConnection)
				
				if(targetConnectionId != cid) {
					writer.WriteString("Got New Connection \n")	
				} else if (targetConnectionId == cid) {
					writer.WriteString("Welcome to the system.... \n")		
				}
				writer.Flush()
			})
		}	
	}(connectionPool)	
	
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGTSTP, syscall.SIGQUIT )
	
	for {
		<-c
		fmt.Println("killing program")
		break;
	}		
}