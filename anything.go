package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	_ "regexp"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var mode string

func main() {
	modePtr := flag.String("mode", "", "which mode to run")

	portPtr := flag.String("port", "8081", "Which port to run")
	flag.Parse()
	fmt.Println("word:", *modePtr)
	fmt.Println("port:", *portPtr)

	if *modePtr == "dev" {
		mode = "dev"
	}

	router := httprouter.New()
	router.GET("/:anything", Index)
	router.POST("/:anything", Index)
	router.PUT("/:anything", Index)
	router.DELETE("/:anything", Index)

	router.GET("/:anything/:anything", Index)
	router.POST("/:anything/:anything", Index)
	router.PUT("/:anything/:anything", Index)
	router.DELETE("/:anything/:anything", Index)

	router.GET("/:anything/:anything/:anything", Index)
	router.POST("/:anything/:anything/:anything", Index)
	router.PUT("/:anything/:anything/:anything", Index)
	router.DELETE("/:anything/:anything/:anything", Index)

	if os.Getenv("LISTEN_PID") == strconv.Itoa(os.Getpid()) {
		// systemd run
		f := os.NewFile(3, "from systemd")
		l, err := net.FileListener(f)
		if err != nil {
			log.Fatal(err)
		}
		http.Serve(l, nil)
	} else {
		// manual run
		//		log.Fatal(http.ListenAndServe(":8080", nil))
		log.Fatal(http.ListenAndServe(":"+*portPtr, router))
	}
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Print("I am here")
	w.Header().Set("Server", "Anything Server")
	w.WriteHeader(200)
	return
}
