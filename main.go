package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
	"log"

	// "gopkg.in/mgo.v2"
	"github.com/emicklei/go-restful"
	
	// "github.com/gaocegege/the-big-brother-is-watching-you/util"
)

// tbbPort is the port to listen on
var tbbPort = flag.Int("port", 8080, "The port to listen on")

const (
	// MongoGracePeriod is the grace period waiting for mongodb to be running.
	MongoGracePeriod = 30 * time.Second
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	// session, err := mgo.Dial(util.GetStringEnvWithDefault("MONGO_DB_IP", "localhost"))
	// if err != nil {
	// 	log.Fatal("Error dailing mongodb")
	// }

	server := &http.Server{Addr: fmt.Sprintf(":%d", *tbbPort), Handler: restful.DefaultContainer}
	
	log.Printf("Server listening on %d", *tbbPort)
	log.Fatal(server.ListenAndServe())
}
