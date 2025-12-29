package entity

import (
	"time"

	"github.com/oklog/ulid"
)

type EventFeedback struct {
	Id           ulid.ULID `json:"id" db:"id"`
	InvocationId string    `json:"invocation_id" db:"invocation_id"`
	UserId       int64     `json:"user_id" db:"user_id"`
	IsPositive   bool      `json:"is_positive" db:"is_positive"`
	Comments     *string   `json:"comments" db:"comments"`
	ErrorType    *string   `json:"error_type" db:"error_type"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
