package main

// on hold for now
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")
}

type G struct {
	id        int
	status    string
	pc        uintptr
	waitingON interface{}
}

type M struct {
	id     int
	curG   *G
	p      *os.PathError
	locked bool
}

type P struct {
	id     int
	runq   []*G
	gfree  []*G
	mcache *simalloc
}
