package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1325A(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, 1, n-1)
	}
}

//func main() { CF1325A(os.Stdin, os.Stdout) }
