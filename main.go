package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	service "tracker/src"
)

func main() {

	Addr := ":8081"
	s := &http.Server{
		Addr:           Addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}
	file, err := os.OpenFile("logs.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	defer file.Close()

	style := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", style))

	http.HandleFunc("/", service.MainPage)
	http.HandleFunc("/search-filter", service.SearchFilterHandler)
	http.HandleFunc("/artist/", service.ArtistPage)
	fmt.Println("Listening on the Addr " + Addr + "\nhttp://localhost" + Addr + "/")
	log.Println("Listening on the Addr " + Addr)

	go Openbrowser("http://localhost" + Addr)

	if err := s.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//Openbrowser - Функция для открытия браузера
func Openbrowser(zz string) {
	time.Sleep(time.Second)
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", zz).Start()
	case "linux":
		err = exec.Command("xdg-open", zz).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}
