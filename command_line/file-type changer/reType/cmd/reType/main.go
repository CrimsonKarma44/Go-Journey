package main

import (
	"flag"
	"fmt"
	"os"
	"reType"
)

func main() {
	DIR, _ := os.Getwd()
	changer := flag.Bool("c", false, "Input extension types")
	flag.Parse()
	if len(flag.Args()) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	} else if len(flag.Args()) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	oldExt := flag.Arg(0)
	newExt := flag.Arg(1)
	if !*changer {
		fmt.Println("missing tag")
		os.Exit(1)
	}
	newFile := reType.File{
		Directory:    DIR,
		OldExtension: oldExt,
		NewExtension: newExt,
	}
	newFile.PatternFix()
	os.Exit(1)

	//scanner := bufio.NewScanner(os.Stdin)
	//scanner = bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//fmt.Println(DIR)
}
