package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1951D(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		if k != n && k > (n+1)/2 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			Fprintln(out, 2)
			Fprintln(out, n-k+1, 1)
		}
	}
}

//func main() { cf1951D(os.Stdin, os.Stdout) }
