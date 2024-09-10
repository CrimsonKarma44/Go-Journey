package main

import (
	"os"
	"regexp"
)

type File struct {
	Directory    string
	OldExtension string
	NewExtension string
}

func (newFile *File) PatternFix() {
	file := obtainData(newFile.Directory, newFile.OldExtension)
	patt := regexp.MustCompile(`(.+).` + newFile.OldExtension)

	for i := 0; i < len(file); i++ {
		err := os.Rename(file[i], patt.ReplaceAllString(file[i], "${1}."+newFile.NewExtension))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	newFile := File{
		Directory:    "/home/karma/Downloads",
		OldExtension: "HEIC",
		NewExtension: "png",
	}
	newFile.PatternFix()
}

func obtainData(directory, oldExtension string) []string {
	requiredCollection := make([]string, 0)

	entry, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	for _, i2 := range entry {
		if pattern(i2.Name(), oldExtension) {
			requiredCollection = append(requiredCollection, directory+"/"+i2.Name())
		}
	}

	return requiredCollection
}

func pattern(file, oldExtension string) bool {
	patt := `(.+).` + oldExtension
	match, err := regexp.MatchString(patt, file)
	if err == nil {
		if match == true {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
