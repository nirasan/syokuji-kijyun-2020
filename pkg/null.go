package pkg

import "strconv"

type NullFloat64 struct {
	Value float64
	Valid bool
}

func NewNullFloat64(v interface{}) NullFloat64 {
	if vv, ok := v.(float64); ok {
		return NullFloat64{vv, true}
	}
	return NullFloat64{0, false}
}

func NullFloat64FromString(s string) NullFloat64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return NewNullFloat64(nil)
	}
	return NewNullFloat64(v)
}

func (n NullFloat64) ValueOrZero() float64 {
	if n.Valid {
		return n.Value
	}
	return 0
}

func (n NullFloat64) Flatten() (float64, bool) {
	return n.Value, n.Valid
}
