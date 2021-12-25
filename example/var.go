package example

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func StudVar() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Println(*n)
	if !*n {
		fmt.Println()
	}

}
