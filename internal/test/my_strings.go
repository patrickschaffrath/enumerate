package test

type MyString string // @enumerate

type MyStringWithoutConst string // @enumerate

type MyStringz string

type myStringz string

const (
	MyStringThis             MyString = "this"
	MyStringIs, MyStringFine MyString = "is", "fine"
	myStringNope             MyString = "nope"
)

const MyStringToo MyString = "too"

const myStringToo MyString = "too"

const MyStringzNope MyStringz = "nope"
