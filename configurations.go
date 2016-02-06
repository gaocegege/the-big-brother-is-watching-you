package main

import (
	"flag"
	"log"
	"time"

	"github.com/gaocegege/the-big-brother-is-watching-you/api"
	"github.com/gaocegege/the-big-brother-is-watching-you/source"
	"github.com/gaocegege/the-big-brother-is-watching-you/storage"
	"github.com/gaocegege/the-big-brother-is-watching-you/timer"
	"github.com/gaocegege/the-big-brother-is-watching-you/worker"

	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2"
)

var (
	mongoIP = flag.String("mongo-db-ip", "localhost", "the location of the mongodb")
	period  = flag.Int64("poll-period", 60, "poll period of the ticker")
)

// Init the main process
func Init(session *mgo.Session) {
	if !flag.Parsed() {
		flag.Parse()
	}

	// alloc the collection managers
	vendorCM := storage.NewVendorCollectionManager(session)
	recordCM := storage.NewRecordCollectionManager(session)

	// alloc source manager
	sourceM, err := source.NewManager()
	if err != nil {
		log.Fatal(err)
	}

	// alloc the worker
	worker := worker.NewWorker(sourceM, vendorCM, recordCM)

	// alloc timer
	timer, err := timer.NewTimer(time.Duration(*period), worker)
	if err != nil {
		log.Fatal(err)
	}

	timer.Run()

	// run rest service
	runRESTService()
}

// runRESTService runs the rest service
func runRESTService() {
	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON, "text/plain", "text/event-stream").
		Produces(restful.MIME_JSON, "text/plain", "text/event-stream")

	log.Print("Register API to server...")

	// register healthz to rest service
	registerHealthzCheckAPI(ws)

	restful.Add(ws)
}

// getMongoSession returns a mongo session
func getMongoSession() *mgo.Session {
	session, err := mgo.Dial(*mongoIP)
	if err != nil {
		log.Fatal("Error dailing mongodb")
	}
	session.SetMode(mgo.Strong, true)

	return session
}

func registerHealthzCheckAPI(ws *restful.WebService) {
	ws.Route(ws.GET("/healthz").To(api.Healthz).Doc("Health Check"))
}
