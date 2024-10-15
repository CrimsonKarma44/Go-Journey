package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const storagePath = "storage/storage.json"

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return wc
}

func main() {
	lines := flag.Bool("l", false, "Count Lines")
	binary := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *binary))
}
