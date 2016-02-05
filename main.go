package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

// tbbPort is the port to listen on
var (
	tbbPort = flag.Int("port", 8080, "The port to listen on")
)

const (
	// MongoGracePeriod is the grace period waiting for mongodb to be running.
	MongoGracePeriod = 30 * time.Second
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	server := &http.Server{Addr: fmt.Sprintf(":%d", *tbbPort), Handler: restful.DefaultContainer}

	log.Printf("Server listening on %d", *tbbPort)
	log.Fatal(server.ListenAndServe())
}
