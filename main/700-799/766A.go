package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF766A(in io.Reader, out io.Writer) {
	var s, t string
	if Fscan(bufio.NewReader(in), &s, &t); s == t {
		Fprint(out, -1)
	} else if len(s) > len(t) {
		Fprint(out, len(s))
	} else {
		Fprint(out, len(t))
	}
}

//func main() { CF766A(os.Stdin, os.Stdout) }
