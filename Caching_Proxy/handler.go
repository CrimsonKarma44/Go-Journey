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
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	fmt.Println("X-Cache:", res.Header.Get("X-Cache"))
	w.Header().Set("X-Cache", res.Header.Get("X-Cache"))
	w.Write([]byte(body))
	fmt.Println("Request made...")
}
