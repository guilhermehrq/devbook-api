package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		ErrorMessage string `json:"errorMessage"`
	}{
		ErrorMessage: err.Error(),
	})
}
