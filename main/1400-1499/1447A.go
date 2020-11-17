package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1447A(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, n)
		for i := 1; i <= n; i++ {
			Fprint(out, i, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1447A(os.Stdin, os.Stdout) }
