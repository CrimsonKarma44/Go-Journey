package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateViewAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("Retrieving Posts ...")
		text := `{"Love":4}`
		jsoned, err := json.Marshal(text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write([]byte(jsoned))
	}
	if r.Method == http.MethodPost {
		fmt.Println("Creating Post ...")
	}
}
