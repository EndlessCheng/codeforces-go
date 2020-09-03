package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF959A(in io.Reader, out io.Writer) {
	var n int
	if Fscan(in, &n); n&1 > 0 {
		Fprint(out, "Ehab")
	} else {
		Fprint(out, "Mahmoud")
	}
}

//func main() { CF959A(os.Stdin, os.Stdout) }
