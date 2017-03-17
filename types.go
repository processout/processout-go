package processout

import "time"

// NilString returns the pointer to the string value
func NilString(s string) *string {
	return &s
}

// NilInt64 returns the pointer to the integer value
func NilInt64(i int64) *int64 {
	return &i
}

// NilFloat64 returns the pointer to the float value
func NilFloat64(f float64) *float64 {
	return &f
}

// NilBool returns the pointer to the bool value
func NilBool(b bool) *bool {
	return &b
}

// NilTime returns the pointer to the time value
func NilTime(t time.Time) *time.Time {
	return &t
}
