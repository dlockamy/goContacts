package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/dlockamy/goContacts/app"
	"github.com/dlockamy/goContacts/controllers"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)


	router.HandleFunc("/api/user/new",
		controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login",
		controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/me/contacts",
		controllers.GetContactsFor).Methods("GET")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}