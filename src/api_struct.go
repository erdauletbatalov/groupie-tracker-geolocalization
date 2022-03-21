package service

import "sync"

// Artist is
type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	DatesLocations map[string][]string
}

// City is
type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// Relation is
type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// General is
type General struct {
	Artists []Artist
	Cities  []City
}

// AllStruct is
type AllStruct struct {
	General    General
	Search     General
	Filter     General
	Open       bool
	Noncorrect bool
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}
