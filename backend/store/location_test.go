package store

import (
	"bz.moh.epi/godatatools/models"
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed locs2.json
var locs []byte

func TestRead(t *testing.T) {
	locations := make(map[string]models.AddressLocation)
	err := json.Unmarshal(locs, &locations)
	if err != nil {
		t.Errorf("failed to unmarshal locations: %v", err)
	}
	loc := locations["1e16e36b-f2ff-4265-8a1f-9a8ea650c719"]
	t.Logf("retrieved location: %v", loc)
}
