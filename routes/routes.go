package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/guilherme8/go-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.FindPersonalityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
