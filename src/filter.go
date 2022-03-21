package service

import (
	"strconv"
)

// MaxInt is Maximum int64 9223372036854775807
const MaxInt = int(^uint(0) >> 1)

// Filter Filters Artists by given inputs
func Filter(input Inputs) []Artist {
	fromCD, err := strconv.Atoi(input.CD.From)
	if err != nil {
		fromCD = 0
	}
	toCD, err := strconv.Atoi(input.CD.To)
	if err != nil {
		toCD = MaxInt
	}
	fromFAD, err := strconv.Atoi(input.FAD.From)
	if err != nil {
		fromFAD = 0
	}
	toFAD, err := strconv.Atoi(input.FAD.To)
	if err != nil {
		toFAD = MaxInt
	}
	fromNOM, err := strconv.Atoi(input.NOM.From)
	if err != nil {
		fromNOM = 0
	}
	toNOM, err := strconv.Atoi(input.NOM.To)
	if err != nil {
		toNOM = MaxInt
	}
	Artists := []Artist{}
	for _, art := range API.General.Artists {
		if input.Chechboxes.CDCheck != "" {
			if !cdToFrom(fromCD, toCD, art) {
				continue
			}
		}
		if input.Chechboxes.FADCheck != "" {
			if !fadToFrom(fromFAD, toFAD, art) {
				continue
			}
		}
		if input.Chechboxes.NOMCheck != "" {
			if !nomToFrom(fromNOM, toNOM, art) {
				continue
			}
		}
		if input.Chechboxes.LocationsCheck != "" {
			if !locationFilter(input.Loc, art) {
				continue
			}
		}
		Artists = append(Artists, art)
	}
	return Artists
}

// Check if given Artist's Creation date is between given values
func cdToFrom(from, to int, art Artist) bool {
	if art.CreationDate >= from && art.CreationDate <= to {
		return true
	}
	return false
}

// Check if given Artist's First Album date is between given values
func fadToFrom(from, to int, art Artist) bool {
	compare, _ := strconv.Atoi(art.FirstAlbum[len(art.FirstAlbum)-4:])
	if compare >= from && compare <= to {
		return true
	}
	return false
}

// Check if given Artist's Number of members is between given values
func nomToFrom(from, to int, art Artist) bool {
	if len(art.Members) >= from && len(art.Members) <= to {
		return true
	}
	return false
}

// Check if given Artist has checked values
func locationFilter(location []string, art Artist) bool {
	for _, artGen := range API.General.Artists {
		if art.ID == artGen.ID {
			for _, city := range location {
				if _, have := artGen.DatesLocations[city]; !have {
					return false
				}
			}
			return true
		}
	}
	return false
}
