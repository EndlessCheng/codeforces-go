package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF493D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	if n&1 > 0 {
		Fprint(out, "black")
	} else {
		Fprint(out, "white\n1 2")
	}
}

//func main() { CF493D(os.Stdin, os.Stdout) }
