package reType

import (
	"os"
	"regexp"
)

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
