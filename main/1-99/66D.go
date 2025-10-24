package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf66D(in io.Reader, out io.Writer) {
	n := 0
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, -1)
		return
	}
	Fprintln(out, 10)
	Fprintln(out, 15)
	for i := range n - 2 {
		Fprintln(out, (9+i)*6)
	}
}

//func main() { cf66D(os.Stdin, os.Stdout) }
