package Caching_Proxy

import "net/http"

func Middleware(next func(http.ResponseWriter, *http.Request, string, *http.Client), origin string, client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r, origin, client)
	}
}
