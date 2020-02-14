package main

import(
	//Native
	"fmt"
	"flag"
	"net/http"
	"net/http/httputil"
	
	//3rd Party Packages
	"github.com/google/uuid"
	
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	
	//Out Custom Packages		
	"github.com/pr0con/go_public_modules/procon_data"
	"github.com/pr0con/go_public_modules/procon_asyncq"
	"github.com/pr0con/go_public_modules/procon_redis"
)

var pool = procon_data.NewPool()
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
	
	id, err := uuid.NewRandom() 
	if err != nil { fmt.Println(err) }
	
	c.Uuid = "ws-"+id.String()	
	procon_data.SendMsg("^vAr^", "client-websocket-id", c.Uuid, c);
	
	
	//create a fat client add fat client to pool
	fat_client := &procon_data.FatClient{
		Id: c.Uuid,
        Conn: c,
        Pool: pool,
    }
	pool.Register <- fat_client	
	
	//add key to redis
	tobj := procon_redis.NewRedisTask("set-key", c.Uuid, "noop")
	procon_asyncq.TaskQueue <- tobj		
	
	Loop:
		for {
			in := procon_data.Msg{}	
		
			err := c.ReadJSON(&in)
			if err != nil {
				fmt.Println("Error @c.ReadJSON: ", err)
				pool.Unregister <- fat_client
					
				tobj := procon_redis.NewRedisTask("del-key", c.Uuid, "noop")
				procon_asyncq.TaskQueue <- tobj				
				
				c.Close()
				break Loop
			}
			
			switch(in.Type) {
				case "client-id-ack":
					fmt.Println(in.Data)
				default:
					break;
			}
		}		
}

func main() {
	go pool.Start()	
	go procon_asyncq.StartTaskDispatcher(9)
	
	r := mux.NewRouter()	
	r.HandleFunc("/ws", handleC)
	
	fmt.Println("Server running on port 0.0.0.0:1200")
	http.ListenAndServeTLS(*addr,"/etc/letsencrypt/live/xbin.pr0con.com/cert.pem", "/etc/letsencrypt/live/xbin.pr0con.com/privkey.pem", r)		
}