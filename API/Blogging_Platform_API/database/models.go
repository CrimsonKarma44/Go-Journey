// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package Blogging_Platform_API

import (
	"encoding/json"
	"time"
)

type Blog struct {
	ID        int64
	Title     string
	Content   string
	Category  string
	Tags      json.RawMessage
	Updatedat time.Time
	Createdat time.Time
}
