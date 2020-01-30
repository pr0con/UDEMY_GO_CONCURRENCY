package main

import(
	"fmt"
	"sync"
	"time"
)


/*    
		type Map
        func (m *Map) Delete(key interface{})
        func (m *Map) Load(key interface{}) (value interface{}, ok bool)
        func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
        func (m *Map) Range(f func(key, value interface{}) bool)
        func (m *Map) Store(key, value interface{})
*/


func main() {
	m := sync.Map{}
		
	m.Store("Key 1", "Value 1")
	m.Store("Key 2", "value 2")
	m.Store("Key 3", "value 3")
	m.Store("Key 4", "value 4")
	m.Store("Key 5", "value 5")	
	
	<-time.After(5 * time.Millisecond)
	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v -> value: %v\n", key, value)
		return true
	})	
	
	LOS1, ok := m.LoadOrStore("Key 7", "Value 7")
	if ok { fmt.Println("Created or Loaded") }
	fmt.Printf("Key 7: %v \n", LOS1)
	<-time.After(5 * time.Millisecond)
	
	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v -> value: %v\n", key, value)
		<-time.After(5 * time.Millisecond)
		return true
	})	
	
	m.Delete("Key 5")
	<-time.After(5 * time.Millisecond)	
	
	fmt.Println(" ")
	
	m.Range(func(key interface{}, value interface{}) bool {
		fmt.Printf("key: %v -> value: %v\n", key, value)
		<-time.After(5 * time.Millisecond)
		return true
	})
	
	LOS2, ok := m.Load("Key 4")
	if ok { fmt.Println(LOS2) }
}