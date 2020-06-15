package main

import (
	"github.com/GabielFemi/tammyNal/tammyNal"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	// serve assets directory
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.HandleFunc("/", tammyNal.Index)

	server := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: router,
	}

	logrus.Println("Listening on ", server.Addr)
	logrus.Fatalln(server.ListenAndServe())
}