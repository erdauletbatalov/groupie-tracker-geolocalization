package service

import (
	"errors"
	"strconv"
	"strings"
)

// SearchArtist searches for people who have matches with the search word in the "Search by" list
func (FilteredArtists *General) SearchArtist(search, sort string) ([]Artist, error) {
	res := []Artist{}
	switch sort {
	case "all":
		{
			a, atoierr := strconv.Atoi(search)
			if atoierr == nil {
				for _, art := range FilteredArtists.Artists {
					if art.CreationDate == a {
						res = append(res, art)
						continue
					}
				}
				if len(res) == 0 {
					return res, errors.New("empty list")
				}
			}
			for _, art := range FilteredArtists.Artists {
				if strings.Contains(art.Name, search) {
					res = append(res, art)
					continue
				}

				if strings.Contains(art.FirstAlbum, search) {
					res = append(res, art)
					continue
				}
				for j := 0; j < len(art.Members); j++ {
					if strings.Contains(art.Members[j], search) {
						res = append(res, art)
						break
					}
				}
				if IsDuplication(res, art.ID) {
					continue
				}
				for _, artFlt := range FilteredArtists.Artists {
					if art.ID == artFlt.ID {
						if _, have := artFlt.DatesLocations[search]; have {
							res = append(res, art)
							continue
						}
						for key := range artFlt.DatesLocations {
							if strings.Contains(key, search) {
								res = append(res, art)
								break
							}
						}
					}
				}
			}
			if len(res) == 0 {
				return res, errors.New("empty list")
			}
			return res, nil
		}
	case "artist":
		{
			for _, art := range FilteredArtists.Artists {
				if strings.Contains(art.Name, search) {
					res = append(res, art)
					continue
				}
			}
			if len(res) == 0 {
				return res, errors.New("empty list")
			}
			return res, nil
		}
	case "members":
		{
			for _, art := range FilteredArtists.Artists {
				for j := 0; j < len(art.Members); j++ {
					if strings.Contains(art.Members[j], search) {
						res = append(res, art)
						break
					}
				}
			}
			if len(res) == 0 {
				return res, errors.New("empty list")
			}
			return res, nil
		}
	case "creation":
		{
			a, atoierr := strconv.Atoi(search)
			if atoierr != nil {
				return res, errors.New("not a number")
			}
			for _, art := range FilteredArtists.Artists {
				if art.CreationDate == a {
					res = append(res, art)
					continue
				}
			}
			if len(res) == 0 {
				return res, errors.New("empty list")
			}
			return res, nil
		}
	case "album":
		{
			for _, art := range FilteredArtists.Artists {
				if strings.Contains(art.FirstAlbum, search) {
					res = append(res, art)
					continue
				}
			}
			if len(res) == 0 {
				return res, errors.New("empty list")
			}
			return res, nil
		}
	case "location":
		{
			for _, art := range FilteredArtists.Artists {
				for key := range art.DatesLocations {
					if strings.Contains(key, search) {
						res = append(res, art)
						break
					}
				}
			}
			if len(res) == 0 {
				return res, errors.New("empty list")
			}
			return res, nil
		}
	default:
		{
			return res, errors.New("empty list")
		}
	}
}

//IsDuplication - возвращает истину, если в списке артистов есть дупликат
func IsDuplication(art []Artist, ID int) bool {
	for _, val := range art {
		if val.ID == ID {
			return true
		}
	}
	return false
}
