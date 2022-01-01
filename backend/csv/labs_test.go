package csv

import (
	"bz.moh.epi/godatatools/store"
	"context"
	"encoding/csv"
	"os"
	"testing"
	"time"
)

func TestWriteLabs(t *testing.T) {
	database := os.Getenv("MONGO_DB")
	uri := os.Getenv("MONGO_URI")
	dbStore, err := store.New(uri, database)
	if err != nil {
		t.Fatalf("failed to create the mongo client: %v", err)
	}
	ctx := context.Background()
	//connect
	if err := dbStore.Connect(ctx); err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}
	defer dbStore.Disconnect(ctx)

	startDate, _ := time.Parse(layoutISO, "2021-12-22")
	endDate, _ := time.Parse(layoutISO, "2021-12-31")
	labTests, err := dbStore.FindLabTestsByDateRange(ctx, &startDate, &endDate)
	if err != nil {
		t.Fatalf("Fetching data failed: %v", err)
	}
	f, err := os.Create("labs.csv")
	defer f.Close()
	csvWriter := csv.NewWriter(f)
	if err := WriteLabs(csvWriter, labTests); err != nil {
		t.Fatalf("failed to write lab csv %v", err)
	}
	csvWriter.Flush()
}
