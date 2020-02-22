package welcomeHttpHandlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	data := [1]string{"Hello"}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshaling output failed:", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(jsonBytes); err != nil {
		log.Println("Writing response failed:", err)
	}
}