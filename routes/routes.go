package routes

import (
	"log"
	"net/http"

	"github.com/arthur-rc18/Api-Go-Rest/controllers"
	"github.com/gorilla/mux"
)

// With this function,it is registered a new route with a matcher for the URL path
func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalitys", controllers.AllPersonalitys).Methods("Get")
	r.HandleFunc("/api/personalitys/{id}", controllers.ReturnPersonality).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
