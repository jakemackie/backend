package routes

import (
	"log"
	"net/http"

	"server/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000"},
		AllowedMethods:   []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowCredentials: true,
		Debug:            true,
	})

    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/api/videos", controllers.AllVideos).Methods("GET")
    myRouter.HandleFunc("/api/videos", controllers.PostVideo).Methods("POST")

	handler := c.Handler(myRouter)

    log.Fatal(http.ListenAndServe(":8000", handler))
}