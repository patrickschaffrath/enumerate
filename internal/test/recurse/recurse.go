package recurse

type (
	MyUint uint64 // @enumerate
)

type (
	ignore    uint64 // @enumerate
	ignoreToo uint64
)

type IgnoreTooo uint64

const (
	MyUintThis           MyUint = 1
	MyUintIs, MyUintFine MyUint = 2, 3
)

const MyUintToo MyUint = 4

const ignoreMe ignore = 1
