package null

import "strconv"

func StrToNilBool(val string) *bool {
	var b *bool
	if s, err := strconv.ParseBool(val); err == nil {
		b = &s
	}
	return b
}

func PtrBool(b bool) *bool {
	return &b
}

func StrToNilInt(val string) *int {
	var i *int
	if s, err := strconv.Atoi(val); err == nil {
		i = &s
	}
	return i
}
