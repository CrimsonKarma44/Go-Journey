package main

import (
	"Broadcast_Server"
	"flag"
)

func main() {
	start := flag.Bool("start", false, "start server")
	connect := flag.Bool("connect", false, "connect server")

	flag.Parse()

	switch {
	case *start:
		Broadcast_Server.Url()
	case *connect:
		Broadcast_Server.Connect()

	}
}