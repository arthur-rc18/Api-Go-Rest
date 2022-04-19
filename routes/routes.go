package routes

import (
	"log"
	"net/http"

	"github.com/arthur-rc18/Api-Go-Rest/controllers"
	"github.com/arthur-rc18/Api-Go-Rest/middleware"
	"github.com/gorilla/mux"
)

// With this function,it is registered a new route with a matcher for the URL path
func HandleRequest() {

	// NewRouter returns a new router instance
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalitys", controllers.AllPersonalitys).Methods("Get")
	r.HandleFunc("/api/personalitys/{id}", controllers.ReturnPersonality).Methods("Get")    // The Get method will receive the data
	r.HandleFunc("/api/personalitys", controllers.CreateNewPersonality).Methods("Post")     // The Post method will post the data
	r.HandleFunc("/api/personalitys/{id}", controllers.DeletePersonality).Methods("Delete") // The Delete method will delete the specified data
	r.HandleFunc("/api/personalitys/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
