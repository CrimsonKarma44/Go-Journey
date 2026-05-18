package Caching_Proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, alias string, client *http.Client) {
	fmt.Println("Sending proxy request...")
	res, err := client.Get(alias + r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Debug: Print headers from upstream
	fmt.Printf("Upstream Headers: %v\n", res.Header)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error reading body:", err)
		return
	}

	xCache := res.Header.Get("X-Cache")
	fmt.Println("X-Cache:", xCache)
	w.Header().Set("X-Cache", xCache)
	w.Write([]byte(body))
	fmt.Println("Request made...")
}
