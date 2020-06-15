package main

import (
	"github.com/GabielFemi/tammyNal/tammyNal"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter
	router().HandleFunc("/", tammyNal.Index)
}