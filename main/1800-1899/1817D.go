package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1817D(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	if k < n/2 {
		Fprint(out, strings.Repeat("LDRU", k-1), "L")
	} else {
		Fprint(out, strings.Repeat("RDLU", n-k-2), strings.Repeat("LDLU", n), "RDL")
	}
}

//func main() { cf1817D(os.Stdin, os.Stdout) }
