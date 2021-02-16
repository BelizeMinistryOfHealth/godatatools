package main

import (
	"bz.moh.epi/godatatools"
	"context"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"log"
	"os"
)

func main() {
	port := "8080"
	// Use PORT env variable, or default to 8080
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	godatatools.GetServer().BackendBaseURL = fmt.Sprintf("http://localhost:%s", port)

	ctx := context.Background()

	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/cases", godatatools.HandlerCasesByOutbreak); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
