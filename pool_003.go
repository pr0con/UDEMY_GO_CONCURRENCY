
package main

import (
	"net"
	"sync"
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



