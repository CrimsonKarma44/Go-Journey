package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	letterStore = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	fmt.Println("Rangoli")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rangoliLetter := scanner.Text()

	strValue, _ := strconv.Atoi(rangoliLetter)

	for _, i := range RangoliFunc(strValue) {
		fmt.Println(i)
	}
}

// RangoliFunc construct the rangoli pattern
func RangoliFunc(valueLength int) []string {
	value := strings.Split(letterStore, "")[0:valueLength]

	var store []string

	for i := valueLength - 1; i >= 0; i-- {
		var temp []string
		for v := valueLength - 1; v >= i; v-- {
			if v-i > 0 {
				temp = append(temp, value[v])
			} else {
				temp = append(temp, value[i])
			}
		}

		for v := i + 1; v < valueLength; v++ {
			temp = append(temp, value[v])
		}
		store = append(store, strings.Join(dasher(valueLength, temp), ""))
	}

	for i := 1; i < valueLength; i++ {
		var temp []string
		for v := valueLength - 1; v >= i; v-- {
			if v-i > 0 {
				temp = append(temp, value[v])
			} else {
				temp = append(temp, value[i])
			}
		}
		for v := i + 1; v < valueLength; v++ {
			temp = append(temp, value[v])
		}
		store = append(store, strings.Join(dasher(valueLength, temp), ""))
	}

	return store
}

// for embedding dashes in the line
func dasher(length int, storeValue []string) []string {
	var dash []string
	totallength := ((length + length - 1) * 2) - 1
	centerPoxMain := (totallength / 2)
	centerPoxStore := (len(storeValue) / 2)

	for i := 0; i < totallength; i++ {
		dash = append(dash, "-")
	}

	if len(storeValue) == 1 {
		dash[centerPoxMain] = storeValue[0]
	} else {
		deviation := 0
		for i := centerPoxStore; i >= 0; i-- {
			dash[centerPoxMain-deviation] = storeValue[i]
			dash[centerPoxMain+deviation] = storeValue[i]
			deviation += 2
		}
	}

	return dash
}
