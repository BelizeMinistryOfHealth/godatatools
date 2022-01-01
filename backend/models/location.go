package models

type AddressLocation struct {
	Id          string `json:"_id"`
	Active      bool   `json:"active"`
	GeoLocation struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"geoLocation"`
	GeographicalLevelId string `json:"geographicalLevelId"`
	Name                string `json:"name"`
	District            string `json:"district"`
}
