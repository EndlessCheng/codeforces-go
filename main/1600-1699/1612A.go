package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1612A(in io.Reader, out io.Writer) {
	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		if x&1 != y&1 {
			Fprintln(out, -1, -1)
		} else {
			Fprintln(out, (x+1)/2, y/2)
		}
	}
}

//func main() { CF1612A(os.Stdin, os.Stdout) }
