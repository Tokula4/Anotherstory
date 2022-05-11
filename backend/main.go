package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brianloveswords/airtable"
	"github.com/ej-limited/auditions/handlers"
	"github.com/ej-limited/auditions/pkg/mail"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	mailuser = os.Getenv("MAIL_USER")
	mailpass = os.Getenv("MAIL_PASS")
	apiKey   = os.Getenv("API_KEY")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := mux.NewRouter()
	fmt.Println(apiKey)
	mc := mail.NewMailClient(mailuser, mailpass)
	c := airtable.Client{
		APIKey: apiKey,
		BaseID: "appk0cwYJZsoXanmK",
	}

	if err != nil {
		log.Println("unable to create airtable client " + err.Error())
		os.Exit(1)
	}
	aHandler := handlers.NewAuditionHandler(mc, &c)
	r.HandleFunc("/register", aHandler.SignUP).Methods("POST")

	log.Println("starting server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}