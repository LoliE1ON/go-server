package main

import (
	"fmt"
	"log"
	"net/http"

	welcomeHttpHandlers "github.com/LoliE1ON/go-server/httpHandlers/welcome"
	"github.com/gorilla/mux"
)

type ServerConfig struct {
	Port int
}

func main() {

	var config = ServerConfig{
		Port: 3500,
	}

	router := mux.NewRouter()
	router.Use(accessControlMiddleware)
	router.HandleFunc("/", welcomeHttpHandlers.Welcome).Methods("GET")

	log.Println("Server started at port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router); err != nil {
		log.Println("Starting server failed: ", err)
	}

}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
