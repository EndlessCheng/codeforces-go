package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1237E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	n *= 3
	if n&(n+5) < 5 {
		Fprint(out, 1)
	} else {
		Fprint(out, 0)
	}
}

//func main() { cf1237E(os.Stdin, os.Stdout) }
