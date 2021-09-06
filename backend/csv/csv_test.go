package csv

import (
	"bz.moh.epi/godatatools/store"
	"context"
	"encoding/csv"
	"os"
	"testing"
)

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
	defer store.Disconnect(ctx)

	cases, _ := store.FindCasesByOutbreak(ctx, outbreakId)

	f, err := os.Create("users.csv")
	defer f.Close()
	//b := &bytes.Buffer{}
	csvWriter := csv.NewWriter(f)
	WriteCases(csvWriter, cases)
	csvWriter.Flush()
}
