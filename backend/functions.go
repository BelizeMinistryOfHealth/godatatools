package godatatools

import (
	"bz.moh.epi/godatatools/auth"
	"bz.moh.epi/godatatools/store"
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var server Server

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	uri := os.Getenv("MONGO_URI")
	database := os.Getenv("MONGO_DB")
	mongoClient, err := store.New(uri, database)
	if err != nil {
		log.Fatalf("could not instantiate the mongo client: %v", err)
	}
	backendBaseURL := "https://us-east1-epi-belize.cloudfunctions.net"
	godataBaseURL := "https://godata-dev.epi.openstep.bz"
	godata := auth.GoData{BaseURL: godataBaseURL}
	ctx := context.Background()
	server = Server{DbRepository: mongoClient, BackendBaseURL: backendBaseURL, GoData: godata}
	if err := server.DbRepository.Connect(ctx); err != nil {
		log.Fatalf("could not connect to mongo: %v", err)
	}
}

func HandlerCasesByOutbreak(w http.ResponseWriter, r *http.Request) {
	corsMid := NewChain(EnableCors(), FileTransferType())
	//if err := server.DbRepository.Connect(r.Context()); err != nil {
	//	log.Fatalf("could not connect to mongo: %v", err)
	//}
	//defer server.DbRepository.Disconnect(r.Context())
	corsMid.Then(server.CasesByOutbreak)(w, r)
}

func HandlerOutbreaks(w http.ResponseWriter, r *http.Request) {
	corsMid := NewChain(EnableCors(), JsonContentType())
	//if err := server.DbRepository.Connect(r.Context()); err != nil {
	//	log.Fatalf("could not connect to mongo: %v", err)
	//}
	//defer server.DbRepository.Disconnect(r.Context())
	corsMid.Then(server.AllOutbreaks)(w, r)
}

func HandlerGoDataAuth(w http.ResponseWriter, r *http.Request) {
	corsMid := NewChain(EnableCors(), JsonContentType())
	corsMid.Then(server.AuthWithGodata)(w, r)
}

func HandlerLabTestSearchResult(w http.ResponseWriter, r *http.Request) {
	mid := NewChain(EnableCors(), JsonContentType())
	mid.Then(server.LabTestResults)(w, r)
}

func HandlerLabTestPrinter(w http.ResponseWriter, r *http.Request) {
	mid := NewChain(EnableCors())
	mid.Then(server.LabTestPdfHandler)(w, r)
}

func HandlerLabTestsReportsCsv(w http.ResponseWriter, r *http.Request) {
	mid := NewChain(EnableCors())
	mid.Then(server.LabTestsByDateRange)(w, r)
}

// Server is exposed to modify the Server settings
func GetServer() *Server {
	return &server
}

func (s Server) RegisterHandlers() {
	http.HandleFunc("/casesByOutbreak", HandlerCasesByOutbreak)
	http.HandleFunc("/auth", HandlerGoDataAuth)
	http.HandleFunc("/outbreaks", HandlerOutbreaks)
	http.HandleFunc("/labTests/searchByName", HandlerLabTestSearchResult)
	http.HandleFunc("/labtest/pdf", HandlerLabTestPrinter)
	http.HandleFunc("/labtestreport/csv", HandlerLabTestsReportsCsv)
}
