package main

import (
	"fmt"
	"time"
)

type user struct {
	name        string    `verify:"min=4,max=10" `
	dateOfBirth time.Time `verify:"date=2006-01-01 00:00:00"`
}

func main() {
	t := user{
		"Vincent",
		func() time.Time {
			value, _ := time.Parse("2006-01-02 15:04:05", "2001-01-23 00:00:00")
			return value
		}(),
	}
	val := time.Now().Sub(t.dateOfBirth)
	val2 := time.Now().Add(-val)
	fmt.Println(val)
	fmt.Println(val2)
	fmt.Println(int64(time.Since(t.dateOfBirth) / (time.Hour * 24 * 365)))
	fmt.Println(resolver(t))
}
