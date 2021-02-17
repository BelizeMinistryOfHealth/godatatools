package store

import (
	"bz.moh.epi/godatatools/models"
	"context"
	"fmt"
	"os"
	"testing"
	"time"
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

	var hospitalizationCases []models.Case
	for _, c := range cases {
		if c.Hospitalizations != nil && len(c.Hospitalizations) > 0 {
			hospitalizationCases = append(hospitalizationCases, c)
		}
	}
	//t.Logf("hospitalizations: %v", hospitalizationCases)
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

func TestStore_OutbreakById(t *testing.T) {
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
	outbreakId := "d54c3aa5-7f43-4733-a482-32d4f8d0b8c4"

	outbreak, err := store.OutbreakById(ctx, outbreakId)
	if err != nil {
		t.Fatalf("OutbreakById() failed: %v", err)
	}

	if outbreak.Name == "" {
		t.Errorf("OutbreakById() name should not be empty")
	}
}

func TestStore_FindCasesByReportingDate(t *testing.T) {
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
	outbreakId := "d54c3aa5-7f43-4733-a482-32d4f8d0b8c4"
	reportingDate, _ := time.Parse("2006-01-02", "2021-02-16")

	cases, err := store.FindCasesByReportingDate(ctx, outbreakId, reportingDate)
	if err != nil {
		t.Fatalf("FindCasesByReportingDate() failed: %v", err)
	}

	if len(cases) == 0 {
		t.Errorf("FindCasesByReportingDate() cases should not be empty")
	}
}

func TestStore_LabTestsForCases(t *testing.T) {
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

	caseIds := []string{"aad8daf1-272d-4ec9-b7bd-7793d5196452", "c0c60423-357d-4407-ba0a-e075efd5d6ae", "3ebc725b-dfc5-4c88-9bce-7362f42725c2"}

	cases, err := store.LabTestsForCases(ctx, caseIds)
	if err != nil {
		t.Fatalf("LabTestsForCases() failed: %v", err)
	}

	if len(cases) != 2 {
		t.Errorf("LabTestsForCases() expected result size to be 2")
	}
}
