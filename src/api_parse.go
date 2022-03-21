package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var doNotUpdateCities bool = false
var enough = 0

// GetAPI - Merges json with our structures every 3 minutes
// FillArtistsWithDatesLocations
// FillGeneralWithCities
// FillCitiesLatLng
func (c *SafeCounter) GetAPI() {
	c.mu.Lock()
	for {
		defer c.mu.Unlock()
		Relation := Relation{}
		artist := GetRequest("https://groupietrackers.herokuapp.com/api/artists")
		relation := GetRequest("https://groupietrackers.herokuapp.com/api/relation")
		err := json.Unmarshal(artist, &API.General.Artists)
		if err != nil {
			log.Println(err)
			return
		}
		err = json.Unmarshal(relation, &Relation)

		if err != nil {
			log.Println(err)
			return
		}

		for i := range API.General.Artists {
			FillArtistsWithDatesLocations(&API.General.Artists[i], Relation)
		}

		if !doNotUpdateCities {
			FillGeneralWithCities()
			FillCitiesLatLng()
			log.Println(API.General.Cities)
		}
		doNotUpdateCities = true

		log.Println("Api has been updated.")
		API.Filter = API.General
		time.Sleep(time.Minute * 3)
	}
}

//GetRequest - takes body from url
func GetRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

// FillArtistsWithDatesLocations Fills artists with datesLocations map
func FillArtistsWithDatesLocations(art *Artist, relation Relation) {
	for _, val := range relation.Index {
		if art.ID == val.ID {
			art.DatesLocations = val.DatesLocations
		}
	}
}

// FillGeneralWithCities fills global array Cities with unique cities
func FillGeneralWithCities() {
	for _, art := range API.General.Artists {
		for key := range art.DatesLocations {
			if !Exists(key) {
				city := City{
					Name: key,
				}
				API.General.Cities = append(API.General.Cities, city)
			}
		}
	}
}

// Exists checks if global array Cities have specific city
func Exists(cityStr string) bool {
	for _, city := range API.General.Cities {
		if city.Name == cityStr {
			return true
		}
	}
	return false
}

// FillCitiesLatLng fills API.General.Cities
// with function FillCityLatLng(city) using goroutines
func FillCitiesLatLng() {
	c := SafeCounter{v: make(map[string]int)}
	for i := range API.General.Cities {
		go c.FillCityLatLng(i)
	}
	time.Sleep(time.Second)
}

// FillCityLatLng fills one city
func (c *SafeCounter) FillCityLatLng(i int) {
	c.mu.Lock()
	var err int
	API.General.Cities[i].Latitude, API.General.Cities[i].Longitude, err = GetCityCoordinates(API.General.Cities[i].Name)
	if err == 0 && enough < 20 {
		enough++
		c.FillCityLatLng(i)
	}
	c.mu.Unlock()
}
