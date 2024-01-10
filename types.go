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

// ToString returns the value of the string pointer, or an empty string
func ToString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

// ToInt64 returns the value of the int64 pointer, or 0
func ToInt64(i *int64) int64 {
	if i == nil {
		return 0
	}

	return *i
}

// ToFloat64 returns the value of the float64 pointer, or 0
func ToFloat64(f *float64) float64 {
	if f == nil {
		return 0
	}

	return *f
}

// ToBool returns the value of the bool pointer, or false
func ToBool(b *bool) bool {
	return b != nil && *b
}

// ToTime returns the value of the time pointer, or an empty time.Time struct
func ToTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}

	return *t
}
