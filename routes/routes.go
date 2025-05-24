package routes

import (
	"log"
	"net/http"

	"github.com/guilherme8/go-api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
