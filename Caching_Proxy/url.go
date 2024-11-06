package Caching_Proxy

import (
	"fmt"
	"log"
	"net/http"
)

func Setup(port string, origin string, client *http.Client) {
	fmt.Println("Starting Proxy Server...")
	http.HandleFunc("/", Middleware(RequestHandler, origin, client))
	fmt.Println("Server Starter ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
