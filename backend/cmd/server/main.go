package main

import (
	"bz.moh.epi/godatatools"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main()  {
	server := godatatools.GetServer()
	server.RegisterHandlers()

	port := os.Getenv("GODATA_TOOLS_PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}

}
