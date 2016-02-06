package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

// tbbPort is the port to listen on
var (
	tbbPort = flag.Int("port", 8080, "The port to listen on")
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	// get the mongodb session
	session := getMongoSession()
	defer session.Close()

	Init(session)

	server := &http.Server{Addr: fmt.Sprintf(":%d", *tbbPort), Handler: restful.DefaultContainer}

	log.Printf("Server listening on %d", *tbbPort)
	log.Fatal(server.ListenAndServe())
}
