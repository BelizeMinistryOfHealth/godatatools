package godatatools

import (
	"bz.moh.epi/godatatools/store"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var server Server

func init() {
	uri := os.Getenv("MONGO_URI")
	database := os.Getenv("MONGO_DB")
	mongoClient, err := store.New(uri, database)
	if err != nil {
		log.Fatalf("could not instantiate the mongo client: %v", err)
	}
	backendBaseURL  := "https://us-east1-epi-belize.cloudfunctions.net"
	server = Server{DbRepository: mongoClient, BackendBaseURL: backendBaseURL}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}


func HandlerCasesByOutbreak(w http.ResponseWriter, r *http.Request) {
	corsMid := NewChain(EnableCors(), JsonContentType())
	if err := server.DbRepository.Connect(r.Context()); err != nil {
		log.Fatalf("could not connect to mongo: %v", err)
	}
	defer server.DbRepository.Disconnect(r.Context())
	corsMid.Then(server.CasesByOutbreak)(w, r)
}

// Server is exposed to modify the Server settings
func GetServer() *Server {
	return &server
}

func (s Server) RegisterHandlers() {
	http.HandleFunc("/casesByOutbreak", HandlerCasesByOutbreak)
}
