package main

import(
	"fmt"
)


type Thing interface {
	actOffTheWall()
}

type Thing1 struct {
	first_name string
}

func (t Thing1) actOffTheWall() {
	fmt.Println(t.first_name, " is being rebellious")
}

type Thing2 struct{
	full_name string	
}
func (t Thing2) actOffTheWall2() {
	fmt.Println(t.full_name, " is being insane")
}

func main() {
	var a Thing1
	var b Thing2
	
	a = Thing1{first_name: "crazy thing one"}
	b = Thing2{full_name: "crazy thing two"} 
	
	a.actOffTheWall()
	b.actOffTheWall2()
	
	var c,d Thing
	c = a
	d = b	
		
	c.actOffTheWall()
	d.actOffTheWall2() 
	
	//error one need to implement the actOffTheWall() method
	//error two isnt defined in the interface, 	
	
	//Interfaces therefor govern/force how struct/objects should behave....
	//Interfaces therefor govern/force how a programs logic should act/operate	
}

