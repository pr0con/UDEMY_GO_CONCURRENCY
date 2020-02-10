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

/**** MODIFIED MUX CLIENT.GO LINE 108 ******
	func (d *Dialer) Dial(urlStr string, requestHeader http.Header) (*Conn, *http.Response, error) {
		+config := &tls.Config{InsecureSkipVerify: true}
		+d.TLSClientConfig = config;
*/
var addr = flag.String("addr", "xbin.pr0con.com:1200", "http service address")


func init() {
	
}

func main() {
	u := url.URL{Scheme: "wss", Host: *addr, Path: "/ws"}
	fmt.Printf("connecting to %s \n", u.String())
	
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil { fmt.Println("Error Dialing Url: %", err) } else {
		
		
		Loop:
			for {
				in := procon_data.Msg{}
				
				err := c.ReadJSON(&in)
				if err != nil {
					fmt.Println("Error @c.ReadJSON: ", err)
					
					c.Close()
					break Loop
				}
				
				switch(in.Type) {
					case "client-websocket-id":
						fmt.Println(in.Data)
						procon_data.SendMsg("^vAr^", "client-id-ack", "Dear Server we got an WsId", c);
					default:
						break;
				}				
				
			}
			
			
	}
}
