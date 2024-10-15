package reType

import (
	"fmt"
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
		fmt.Println(file[i], "changed to", patt.ReplaceAllString(file[i], "${1}."+newFile.NewExtension))
	}
}
