package main

import (
	//Native
	"fmt"
	"flag"
	"net/url"

	//3rd Party
	"github.com/gorilla/websocket"
	
	//Our Modules
	"github.com/pr0con/go_public_modules/procon_data"
)

var addr = flag.String("addr", "xbin.pr0con.com:1200", "http service address")

func init() {
	
}

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	fmt.Printf("connecting to %s \n", u.String())
	
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil { fmt.Println("Error Dialing Url: %", err) } else {
		
		procon_data.SendMsg("noop","hello-world","Hello Server!",c)
		Loop:
			for {
				in := procon_data.Msg{}
				
				err := c.ReadJSON(&in)
				if err != nil {
					fmt.Println("Error @c.ReadJSON: ", err)
					
					c.Close()
					break Loop
				}
				
								
				
			}
			
			
	}
}
