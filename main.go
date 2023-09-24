package main

import (
	"go-web-server-template/api"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port " + ProjectConfig.Port)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", api.RoutesHandler)
	http.ListenAndServe(":"+ProjectConfig.Port, nil)
}
