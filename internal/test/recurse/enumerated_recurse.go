package recurse

import "slices"

var EnumeratedMyUints = []MyUint{
	MyUintThis,
	MyUintIs,
	MyUintFine,
	MyUintToo,
}

func IsMyUint(a any) bool {
	v, ok := a.(MyUint)
	return ok && slices.Contains(EnumeratedMyUints, v)
}
