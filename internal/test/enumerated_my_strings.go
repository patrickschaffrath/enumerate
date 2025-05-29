package test

import "slices"

var EnumeratedMyStrings = []MyString{
	MyStringThis,
	MyStringIs,
	MyStringFine,
}

func IsMyString(a any) bool {
	v, ok := a.(MyString)
	return ok && slices.Contains(EnumeratedMyStrings, v)
}
