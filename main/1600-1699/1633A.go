package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1633A(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n%7 > 0 {
			n -= n % 10
			n += 7 - n%7
		}
		Fprintln(out, n)
	}
}

//func main() { CF1633A(os.Stdin, os.Stdout) }
