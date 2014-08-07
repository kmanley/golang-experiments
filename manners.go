package main

import (
	"fmt"
	"github.com/braintree/manners"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var quitChan = make(chan os.Signal, 1)
var server = manners.NewServer()

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

/* it looks like manners works only if client doesn't use persistent connection;
   e.g. if you run this, then make a request with curl, then ctrl-c this process,
   it shuts down gracefully. But if you make the request with Chrome or Firefox,
   then ctrl-c this process, the http server doesn't quit until you close the
   browser tab */
func waitForSignal() {
	fmt.Println("waiting for signal")
	<-quitChan
	fmt.Println("got signal")
	server.Shutdown <- true
}

func main() {

	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	http.HandleFunc("/foo", Hello)
	go waitForSignal()
	fmt.Println("listening...")
	server.ListenAndServe(":8080", nil)
	fmt.Println("exiting")

}
