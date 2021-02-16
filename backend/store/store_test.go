package store

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestStore_FindCasesByOutbreak(t *testing.T) {
	database := os.Getenv("MONGO_DB")
	uri := os.Getenv("MONGO_URI")
	outbreakId := os.Getenv("OUTBREAK_ID")
	store, err := New(uri, database)
	if err != nil {
		t.Fatalf("failed to create the mongo client: %v", err)
	}
	ctx := context.Background()
	//connect
	if err := store.Connect(ctx); err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}
	defer store.Disconnect(ctx)

	cases, err := store.FindCasesByOutbreak(ctx, outbreakId)
	if err != nil {
		t.Fatalf(fmt.Sprintf("FindCasesByOutbreak(): failed: %v", err))
	}

	if len(cases) == 0 {
		t.Errorf("FindCasesByOutbreak() should return a non-empty list")
	}
}

func TestStore_ListOutbreaks(t *testing.T) {
	database := os.Getenv("MONGO_DB")
	uri := os.Getenv("MONGO_URI")
	store, err := New(uri, database)
	if err != nil {
		t.Fatalf("failed to create the mongo client: %v", err)
	}
	ctx := context.Background()
	//connect
	if err := store.Connect(ctx); err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}
	defer store.Disconnect(ctx)

	outbreaks, err := store.ListOutbreaks(ctx)
	if err != nil {
		t.Fatalf(fmt.Sprintf("ListOutbreaks() failed: %v", err))
	}

	if len(outbreaks) == 0 {
		t.Errorf("ListOutbreaks() should return a non-empty list")
	}

}
