package hash

import (
	"strconv"

	"github.com/spaolacci/murmur3"
)

type customtype interface {
	int | float64 | string
}

type HashFunc[T customtype] struct {
	array        []T
	supportArray []T
}

func (h HashFunc[T]) New() HashFunc[T] {
	return HashFunc[T]{
		array:        make([]T, 10),
		supportArray: make([]T, 0),
	}
}

func (h *HashFunc[T]) Add() {

}

func (h HashFunc[T]) hasher(value T) uint64 {
	var strValue string
	switch v := any(value).(type) {
	case string:
		strValue = v
	case int:
		strValue = strconv.Itoa(v)
	case float64:
		strValue = strconv.FormatFloat(v, 'f', -1, 64)
	default:
		strValue = ""
	}

	hash := murmur3.Sum64([]byte(strValue))
	bucket := hash % 10

	return bucket
}
