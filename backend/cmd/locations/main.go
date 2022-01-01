package main

import (
	"bz.moh.epi/godatatools/models"
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// Exports a json file with location data to an inverted index structure and saves it as json.
// The inverted index makes it faster to retrieve a location by id.
func main() {
	var fileName string
	flag.StringVar(&fileName, "f", "", "Specify the json file with the location data")
	flag.Parse()

	helpMsg := `
	-f : The json file with the location data
	`

	if len(fileName) == 0 {
		log.Errorf("%s", helpMsg)
		os.Exit(-1)
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("readFile failed: %v", err)
		os.Exit(-1)
	}
	var payloadData []models.AddressLocation
	err = json.Unmarshal(content, &payloadData)
	if err != nil {
		log.Errorf("unmarshal failed: %v", err)
		os.Exit(-1)
	}
	var idx = make(map[string]models.AddressLocation)
	for i := range payloadData {
		idx[payloadData[i].Id] = payloadData[i]
	}
	file, err := json.MarshalIndent(idx, "", " ")
	if err != nil {
		log.Errorf("failed to marshal json file: %v", err)
		os.Exit(-1)
	}
	err = ioutil.WriteFile("locations.json", file, 0644)
	if err != nil {
		log.Errorf("failed to write new json file: %v", err)
		os.Exit(-1)
	}
	log.Info("Successfully wrote file: locations.json!!")
	os.Exit(0)
}
