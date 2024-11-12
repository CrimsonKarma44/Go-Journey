package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buffer bytes.Buffer
	bytes.
	buffer.WriteString("love")
	fmt.Println(buffer.String())
	fmt.Println(buffer.Len())
	buffer.Write([]byte(" You"))
	fmt.Println(&buffer)
	buffer.Write([]byte("\nAnd you too"))
	buffer.WriteTo(os.Stdout)
}