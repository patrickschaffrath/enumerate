# Enumerate

Enumerate is a Go tool that collects constants of user defined types and generates slices that contain every exported const for every exported and annotated type.

The tool works on a per file basis, so make sure your constants of a custom type are in the same file as the type declaration.

## Usage

Annotate/comment every exported type you want to enumerate with `@enumerate`, then run

`go run github.com/patrickschaffrath/enumerate@latest`

at your Go project root directory.

## Example

Your project contains `my-project/constants.go` with

```golang
package myproject

type MyString string // @enumerate

const MyStringOne MyString = "one"
const MyStringTwo MyString = "two"
```

After running the tool, it will generate `my-project/enumerated_constants.go` with

```golang
package myproject

import "slices"

var EnumeratedMyStrings = []MyString{
 MyStringOne,
 MyStringTwo,
}

func IsMyString(a any) bool {
 v, ok := a.(MyString)
 return ok && slices.Contains(EnumeratedMyStrings, v)
}
```
