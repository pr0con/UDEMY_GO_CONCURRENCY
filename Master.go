package main

import(
	//Native
	"fmt"
	"flag"
	"net/http"
	"net/http/httputil"
	
	//3rd Party Packages
	//"github.com/google/uuid"
	
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	
	//Out Custom Packages		
	"github.com/pr0con/go_public_modules/procon_data"
)

var addr = flag.String("addr", "0.0.0.0:1200", "http service address")
var upgrader = websocket.Upgrader{} // use default options

func init() {
	
}

func handleC(w http.ResponseWriter, r *http.Request) {
	data, err := httputil.DumpRequest(r, false)
	if err != nil { fmt.Println("Error Handling Http Req.") } else {
		fmt.Printf("%s", string(data))	
	}	
	
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("Error @HandleC Handling Ws Upgrade: ", err)
		return
	}	
	
	Loop:
		for {
			in := procon_data.Msg{}	
		
			err := c.ReadJSON(&in)
			if err != nil {
				fmt.Println("Error @c.ReadJSON: ", err)
				
				c.Close()
				break Loop
			}
			
			fmt.Println(in)
		}		
}

func main() {
	
	r := mux.NewRouter()	
	r.HandleFunc("/ws", handleC)
	
	fmt.Println("Server running on port 0.0.0.0:1200")
	http.ListenAndServe(*addr, r)	
}