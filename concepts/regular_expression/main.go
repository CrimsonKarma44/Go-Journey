package main

import (
	"fmt"
	"regexp"
)

func main() {
	matchString("[A-z]ork", "I love new york city")
	findStringIndex("H[a-z]{4}|[A-z]ork", "Hello guys, welcome to new york city")
	l := regexp.MustCompile("[A-z]ork")
	find := l.MatchString("I love new york city")
	fmt.Println(find)

}

func matchString(s string, pattern string) {
	match, err := regexp.MatchString(pattern, s)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}
}

func findStringIndex(pattern string, s string) {
	patternCompile := regexp.MustCompile(pattern)
	firstMatchIndex := patternCompile.FindStringIndex(s)
	fmt.Println("First matched index", firstMatchIndex[0], "-", firstMatchIndex[1],
		"=", patternCompile.FindString(s))
}
