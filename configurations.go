package main

import (
	"flag"
	"log"
	
	restful "github.com/emicklei/go-restful"
	api "github.com/gaocegege/the-big-brother-is-watching-you/api"
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON, "text/plain", "text/event-stream").
		Produces(restful.MIME_JSON, "text/plain", "text/event-stream")
		
	log.Print("Register API to server...")
	registerHealthzCheckAPI(ws)
	
	restful.Add(ws)
}

func registerHealthzCheckAPI(ws *restful.WebService) {
	ws.Route(ws.GET("/healthz").To(api.Healthz).Doc("Health Check"))
}
