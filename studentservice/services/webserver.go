package services

import (
	"log"
	"net/http"
)

func StartWebserver(port string) {
	router := NewRouter()
	http.Handle("/", router)
	log.Println("Starting the webservice on port ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}
