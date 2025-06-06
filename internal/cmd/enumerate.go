package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/patrickschaffrath/enumerate/internal/enumerator"
)

const version = "v1.0.1"

func Run() {
	flag.Parse()
	if flag.Arg(0) == "version" {
		fmt.Println(version)
		os.Exit(0)
	}
	enumerator.Enumerate()
}
