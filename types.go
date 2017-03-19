package processout

import "time"

// String returns the pointer to the string value
func String(s string) *string {
	return &s
}

// Int64 returns the pointer to the integer value
func Int64(i int64) *int64 {
	return &i
}

// Float64 returns the pointer to the float value
func Float64(f float64) *float64 {
	return &f
}

// Bool returns the pointer to the bool value
func Bool(b bool) *bool {
	return &b
}

// Time returns the pointer to the time value
func Time(t time.Time) *time.Time {
	return &t
}
