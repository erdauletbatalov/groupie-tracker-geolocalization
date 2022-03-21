package service

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//API - Структура, содержащая основную апи и искомых артистов
var API AllStruct
var tpl *template.Template
var err error
var vamlet int = 0

//Info - дефолтная структура, для отображения информации об артисте
type Info struct {
	ArtistID interface{}
	CitiesID interface{}
}

// runs before the main.go
// Parsing all info
func init() {
	c := SafeCounter{v: make(map[string]int)}
	go c.GetAPI()
	tpl, err = template.ParseGlob("static/templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	Relation := Relation{}
	relation := GetRequest("https://groupietrackers.herokuapp.com/api/relation")
	err = json.Unmarshal(relation, &Relation)
	if err != nil {
		log.Println(err)
		return
	}
	for i := range API.General.Artists {
		FillArtistsWithDatesLocations(&API.General.Artists[i], Relation)
	}
}

// MainPage - Handler главной страницы
func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandle(http.StatusNotFound, w, "404 not found")
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandle(http.StatusMethodNotAllowed, w, "405 Status Method Not Allowed")
		return
	}
	Filter := API.General

	AllStruct := API

	AllStruct.Filter = Filter

	if vamlet == 0 {
		AllStruct.Open = false
	} else {
		AllStruct.Open = true
	}
	err = tpl.ExecuteTemplate(w, "index.html", AllStruct)
	vamlet++
	if err != nil {
		err = tpl.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
		if err != nil {
			ErrorHandle(http.StatusInternalServerError, w, err, "500 Internal Server Error")
			return
		}
	}
}

// SearchFilterHandler is a Search and Filter handler. It Filters people by checked checkboxes,
// by typed dates and searches artists if "search" form is not empty
func SearchFilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search-filter" {
		ErrorHandle(http.StatusNotFound, w, "404 not found")
		return
	}
	if r.Method != http.MethodPost {
		ErrorHandle(http.StatusMethodNotAllowed, w, "405 Status Method Not Allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	cities := []string{}
	for _, value := range r.Form["location"] {
		cities = append(cities, value)
	}

	// API.Filter = General{}

	AllStruct := API

	AllStruct.Filter = General{}

	inputs := Inputs{
		CD: CreationDate{
			From: r.FormValue("fromCreation"),
			To:   r.FormValue("toCreation"),
		},
		FAD: FirstAlbumDate{
			From: r.FormValue("fromFAD"),
			To:   r.FormValue("toFAD"),
		},
		NOM: NumberOfMembers{
			From: r.FormValue("fromNOM"),
			To:   r.FormValue("toNOM"),
		},
		Loc: cities,
		Chechboxes: Chechboxes{
			CDCheck:        r.FormValue("CD"),
			FADCheck:       r.FormValue("FAD"),
			NOMCheck:       r.FormValue("NOM"),
			LocationsCheck: r.FormValue("Location"),
		},
	}

	AllStruct.Filter.Artists = Filter(inputs)
	if r.FormValue("FLTR") == "" {
		AllStruct.Filter = API.General
	}

	search := r.FormValue("search")
	if search != "" {
		art, searchErr := AllStruct.Filter.SearchArtist(search, r.FormValue("search_filter"))
		if searchErr != nil {
			AllStruct.Noncorrect = true
		}
		AllStruct.Search = General{Artists: art}
	}
	if vamlet == 0 {
		AllStruct.Open = false
	} else {
		AllStruct.Open = true
	}
	err = tpl.ExecuteTemplate(w, "index.html", AllStruct)
	vamlet++
	if err != nil {
		err = tpl.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
		if err != nil {
			ErrorHandle(http.StatusInternalServerError, w, err, "500 Internal Server Error")
			return
		}
	}
}

// ArtistPage - Handler страницы артиста
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/artist/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		ErrorHandle(http.StatusBadRequest, w, err, "400 Bad request")
		return
	}
	if !(id > 0 && id <= len(API.General.Artists)) {
		ErrorHandle(http.StatusNotFound, w, "404 Not Found")
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandle(http.StatusMethodNotAllowed, w, "405 Method Not Allowed")
		return
	}
	CitiesID := []City{}
	for _, cityLatLng := range API.General.Cities {
		for city := range API.General.Artists[id-1].DatesLocations {
			if cityLatLng.Name == city {
				CitiesID = append(CitiesID, cityLatLng)
				break
			}
		}
	}
	info := &Info{
		ArtistID: API.General.Artists[id-1],
		CitiesID: CitiesID,
	}
	err = tpl.ExecuteTemplate(w, "artistpage.html", info)
	if err != nil {
		err = tpl.ExecuteTemplate(w, "errors.html", http.StatusInternalServerError)
		if err != nil {
			ErrorHandle(http.StatusInternalServerError, w, err, "500 Internal Server Error")
			return
		}
	}
}

// ErrorHandle - Handler страницы ошибки
func ErrorHandle(ErrorStatus int, w http.ResponseWriter, errC ...interface{}) {
	for _, val := range errC {
		log.Println(val)
	}
	w.WriteHeader(ErrorStatus)
	if len(errC) != 2 {
		err = tpl.ExecuteTemplate(w, "errors.html", ErrorStatus)
	} else {
		errC[0] = tpl.ExecuteTemplate(w, "errors.html", ErrorStatus)
	}
	if err != nil {
		fmt.Fprintf(w, "<h1 style=\"text-align: center;\">"+strconv.Itoa(ErrorStatus)+"</h1><br><h3 style=\"text-align: center;\">Internal Server Error</h3><br><a href=\"/\">go back</a>")
	}
}
