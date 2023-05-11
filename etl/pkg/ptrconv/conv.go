package ptrconv

import (
	"time"

	"github.com/google/uuid"
)

// Ptr returns pointer to value
func Ptr[T any](val T) *T {
	return &val
}

// Bool returns pointer to bool
func Bool(val bool) *bool {
	return &val
}

// String returns pointer to string
func String(val string) *string {
	return &val
}

// Time returns pointer to time.Time
func Time(val time.Time) *time.Time {
	return &val
}

// UUID returns pointer to uuid.UUID
func UUID(val uuid.UUID) *uuid.UUID {
	return &val
}

// Int returns pointer to int
func Int(val int) *int {
	return &val
}

// Int32 returns pointer to int32
func Int32(val int32) *int32 {
	return &val
}

// Int64 returns pointer to int64
func Int64(val int64) *int64 {
	return &val
}
