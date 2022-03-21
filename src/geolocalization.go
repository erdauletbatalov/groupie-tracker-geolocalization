package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const helloMessage = "Hello to the weather program. Please enter the name of the city and the weather will show."
const googleAPIURL = "https://maps.googleapis.com/maps/api/geocode/json?key=AIzaSyD3DOfE8-4F5FPEI4kCkroAh5R3K3inYJ0&address="

// GoogleAPIResponse is google api JSON struct
type GoogleAPIResponse struct {
	Results Results `json:"results"`
}

// Results is google api JSON struct lower than GoogleAPIResponse
type Results []Geometry

// Geometry is google api JSON struct lower than Results
type Geometry struct {
	Geometry Location `json:"geometry"`
}

// Location is google api JSON struct lower than Geometry
type Location struct {
	Location Coordinates `json:"location"`
}

// Coordinates is google api JSON struct lower than Location
type Coordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// GetCityCoordinates gets the google api JSON, Converts it to Golang struct,
// Returns only lat, lng and err if json is empty
func GetCityCoordinates(city string) (float64, float64, int) {
	resp, err := http.Get(googleAPIURL + city)

	if err != nil {
		log.Fatal("Fetching google api uri data error: ", err)
		return 0, 0, 1
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("Reading google api data error: ", err)
		return 0, 0, 1
	}

	var data GoogleAPIResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Println(err)
		return 0, 0, 1
	}
	if data.IsEmpty() {
		log.Println("Failed to parse JSON from given city")
		return 0, 0, 0
	}
	lat := data.Results[0].Geometry.Location.Latitude
	lng := data.Results[0].Geometry.Location.Longitude
	log.Println(city + ": Fetching Latitude and longitude ended successful ...")

	return lat, lng, 1
}

// IsEmpty Checks if google api response JSON is empty
func (result *GoogleAPIResponse) IsEmpty() bool {
	if len(result.Results) == 0 {
		return true
	}
	return false
}
