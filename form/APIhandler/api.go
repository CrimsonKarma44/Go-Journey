package APIhandler

import (
	"net/http"
	"os"
)

func allProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		file, err := os.ReadFile("profiles/" + name + ".json")

	}
	return
}
