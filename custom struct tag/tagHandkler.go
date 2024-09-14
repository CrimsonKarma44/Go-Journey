package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func resolver(item interface{}) error {
	for v := 0; v < reflect.ValueOf(item).NumField(); v++ {
		if reflect.TypeOf(item).Field(v).Tag.Get("verify") == "" {
			continue
		}

		listed := strings.Split(reflect.TypeOf(item).Field(v).Tag.Get("verify"), ",")
		for _, i := range listed {
			switch {
			case strings.Contains(i, "min"):
				minimum := toInt(strings.TrimPrefix(i, "min="))
				if len(reflect.ValueOf(item).Field(v).String()) <= minimum {
					return fmt.Errorf("name length is less than required")
				}
			case strings.Contains(i, "max"):
				maximum := toInt(strings.TrimPrefix(i, "max="))
				if len(reflect.ValueOf(item).Field(v).String()) >= maximum {
					return fmt.Errorf("name length is more than required")
				}
			case strings.Contains(i, "date"):
				date := strings.TrimPrefix(i, "date=")
				sample :=
					func() time.Time {
						value, _ := time.Parse("2006-01-02 15:04:05", date)
						return value
					}()

				rs := reflect.ValueOf(item)
				rs2 := reflect.New(rs.Type()).Elem()
				rs2.Set(rs)
				rf := rs2.Field(v)

				f, _ := reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem().Interface().(time.Time)
				if f.After(sample) {
					return fmt.Errorf("%s is below age limit", date)
				}
				//or
				if int64(time.Since(f)/(time.Hour*24*365)) < 18 {
					return fmt.Errorf("%s is below age limit", f)
				}

			}
		}
	}

	return nil
}
func toInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
