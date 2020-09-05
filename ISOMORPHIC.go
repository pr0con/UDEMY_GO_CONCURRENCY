package main

import (
	//Native
	//"os"
	"fmt"
	"flag"
	//"path"
	"net/http"
	"io/ioutil"
	
	
	//3rd Party
	"github.com/gorilla/mux"
)


var addr = flag.String("addr", "0.0.0.0:8000", "http service address")

//Option Two
/*
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HIT NOT FOUND")
	http.ServeFile(w,r, "/var/www/parcel_blueprint/dist")
}

func FileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			NotFound(w,r)
			return
		}
		fsh.ServeHTTP(w,r)
	})
}
*/

//Option Three
type hookedResponseWriter struct {
	http.ResponseWriter
	ignore bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	//Start Insanity Here...
	if status == 404 {
		hrw.ignore = true
		hrw.ResponseWriter.Header().Set("Content-type", "text/html")
		//hrw.ResponseWriter.Header().Set("Cache-Control", "max-age=0")
		hrw.ResponseWriter.WriteHeader(200)
		fmt.Println("Got Here")
		
		data, err := ioutil.ReadFile("/var/www/parcel_blueprint/dist/index.html")
		if err != nil { fmt.Println("Something went Horribly wrong!") } else {
			hrw.ResponseWriter.Write(data)
		}
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.ignore {
		return len(p), nil
	}
	return hrw.ResponseWriter.Write(p)
}

type NotFoundHook struct {
	h http.Handler
}

func (nfh NotFoundHook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nfh.h.ServeHTTP(&hookedResponseWriter{ResponseWriter: w}, r)
}


func main() {
	//Option One && Three
	r := mux.NewRouter()
	
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("/var/www/parcel_blueprint/dist"))) 
	
	r.PathPrefix("/").Handler(NotFoundHook{http.StripPrefix("", http.FileServer(http.Dir("/var/www/parcel_blueprint/dist")))})
	http.ListenAndServeTLS(*addr,"/etc/letsencrypt/live/xbin.pr0con.com/cert.pem", "/etc/letsencrypt/live/xbin.pr0con.com/privkey.pem" , r)
	
	
	//Option Two...
	//http.ListenAndServeTLS(*addr,"/etc/letsencrypt/live/xbin.pr0con.com/cert.pem", "/etc/letsencrypt/live/xbin.pr0con.com/privkey.pem", FileServerWithCustom404(http.Dir("/var/www/parcel_blueprint/dist")))
}