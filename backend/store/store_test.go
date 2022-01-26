package store

import (
	"bz.moh.epi/godatatools/models"
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

const isoLayout string = "2006-01-02"

func TestStore_FindCasesByOutbreak(t *testing.T) {
	outbreakId := os.Getenv("OUTBREAK_ID")
	ctx := context.Background()
	store := setupDb(t, ctx)
	defer store.Disconnect(ctx)

	startDate, _ := time.Parse(isoLayout, "2021-12-21")
	endDate, _ := time.Parse(isoLayout, "2021-12-31")

	cases, err := store.FindCasesByOutbreak(ctx, outbreakId, &startDate, &endDate)
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

func TestStore_LabTestsByCaseName(t *testing.T) {
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

	labTests, err := store.LabTestsByCaseName(ctx, "robErto", "GuerrA")
	if err != nil {
		t.Fatalf("LabTestsByCaseName() failed: %v", err)
	}

	if len(labTests) == 0 {
		t.Errorf("LabTestsByCaseName() expected non-empty list")
	}
}

func TestStore_LabTestById(t *testing.T) {
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

	labTest, err := store.LabTestById(ctx, "a091e02f-bb33-42c9-ac81-699c19ce9790")
	if err != nil {
		t.Fatalf("LabTestsByCaseName() failed: %v", err)
	}

	if len(labTest.ID) == 0 {
		t.Errorf("LabTestsByCaseName() expected non-empty list")
	}
}

func TestStore_FindCasesByPersonIds(t *testing.T) {
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

	ids := []string{"39234b89-05a0-4417-b97f-4fb56e29360f", "1c299f2e-21f5-4164-9e58-88a0540ab75d", "67c31286-1ba7-4f79-a43a-eff76c61dc2d"}
	cases, err := store.FindCasesByPersonIds(ctx, ids)
	if err != nil {
		t.Fatalf("FindCasesByPersonIds() failed: %v", err)
	}

	t.Logf("cases: %v", cases)
}

func TestStore_FindLabTestsByDateRange(t *testing.T) {
	ctx := context.Background()
	store := setupDb(t, ctx)
	defer store.Disconnect(ctx)

	startDate, _ := time.Parse(isoLayout, "2022-01-19")
	endDate, _ := time.Parse(isoLayout, "2022-01-20")

	labTests, err := store.FindLabTestsByDateRange(ctx, &startDate, &endDate)
	if err != nil {
		t.Fatalf("FindLabTestsByDateRange failed: %v", err)
	}
	t.Logf("labTests: %v", labTests)
	var filemon models.LabTestReport

	for i := range labTests {
		if labTests[i].Person.FirstName == "Filemon" {
			filemon = labTests[i]
		}
	}

	t.Logf("filemon: %v", filemon)
}

func setupDb(t *testing.T, ctx context.Context) Store {
	database := os.Getenv("MONGO_DB")
	uri := os.Getenv("MONGO_URI")
	store, err := New(uri, database)
	if err != nil {
		t.Fatalf("failed to create the mongo client: %v", err)
	}

	//connect
	if err := store.Connect(ctx); err != nil {
		t.Fatalf("failed to connect to mongo: %v", err)
	}
	return store
}
