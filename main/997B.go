package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF997B(in io.Reader, out io.Writer) {
	a := [...]int{0, 4, 10, 20, 35, 56, 83, 116, 155, 198, 244}
	var n int64
	Fscan(in, &n)
	if n < 11 {
		Fprint(out, a[n])
	} else {
		Fprint(out, 49*n-247)
	}
}

//func main() { CF997B(os.Stdin, os.Stdout) }
