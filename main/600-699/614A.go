package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf614A(in io.Reader, out io.Writer) {
	var l, r, k int
	Fscan(in, &l, &r, &k)
	ok := false
	for powK := 1; ; powK *= k {
		if powK >= l {
			ok = true
			Fprint(out, " ", powK)
		}
		if powK > r/k {
			break
		}
	}
	if !ok {
		Fprint(out, -1)
	}
}

//func main() { cf614A(os.Stdin, os.Stdout) }
