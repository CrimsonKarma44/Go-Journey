package main

import (
	"Caching_Proxy"
	"fmt"
	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
	"net/http"
	"os"
)

const cacheDir = "/tmp/httpcache"

func main() {
	cache := diskcache.New(cacheDir) // Disk cache
	cacheTransport := httpcache.NewTransport(cache)
	transport := &Caching_Proxy.CacheAwareTransport{Transport: cacheTransport}

	client := &http.Client{
		Transport: transport,
	}
	if len(os.Args) == 5 {
		switch os.Args[1] {
		case "--port":
			switch os.Args[3] {
			case "--origin":

				Caching_Proxy.Setup(os.Args[2], os.Args[4], client)
			default:
				fmt.Println(os.Args[3], "is not a valid input")
			}
		case "--origin":
			switch os.Args[3] {
			case "--port":
				Caching_Proxy.Setup(os.Args[4], os.Args[2], client)
			default:
				fmt.Println(os.Args[3], "is not a valid input")
			}
		default:
			fmt.Println("wrong input")
		}
	} else if len(os.Args) == 2 {
		if os.Args[1] == "--clear-cache" {
			fmt.Println("clearing cache ...")
			err := Caching_Proxy.ClearCache(cacheDir)
			if err != nil {
				fmt.Println("Error clearing cache:", err)
			} else {
				fmt.Println("Cache cleared ...")
			}
		} else {
			fmt.Println("Invalid input")
		}
	} else {
		fmt.Println("Invalid number of arguments")
	}
}
