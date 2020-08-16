package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1266B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		n, k := n%14, n/14
		if k > 0 && 1 <= n && n <= 6 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1266B(os.Stdin, os.Stdout) }
