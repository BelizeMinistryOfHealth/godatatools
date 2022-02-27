package csv

import (
	"bz.moh.epi/godatatools/store"
	"context"
	"encoding/csv"
	"os"
	"testing"
	"time"
)

const isoLayout string = "2006-01-02"

func TestWriteCases(t *testing.T) {
	database := os.Getenv("MONGO_DB")
	uri := os.Getenv("MONGO_URI")
	outbreakId := os.Getenv("OUTBREAK_ID")
	store, err := store.New(uri, database)
	if err != nil {
		t.Fatalf("failed to create the mongo client: %v", err)
	}
	ctx := context.Background()
	//connect
	if err := store.Connect(ctx); err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}
	defer store.Disconnect(ctx) //nolint:errcheck
	startDate, _ := time.Parse(isoLayout, "2021-12-21")
	endDate, _ := time.Parse(isoLayout, "2021-12-31")
	cases, _ := store.FindCasesByOutbreak(ctx, outbreakId, &startDate, &endDate)

	f, err := os.Create("users.csv") //nolint:ineffassign,govet,staticcheck
	defer f.Close()                  //nolint:govet,errcheck,staticcheck
	//b := &bytes.Buffer{}
	csvWriter := csv.NewWriter(f)
	WriteCases(csvWriter, cases) //nolint:errcheck
	csvWriter.Flush()
}
