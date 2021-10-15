package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1581B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, k int
	var n, m int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		if m < n-1 || m > n*(n-1)/2 {
			Fprintln(out, "NO")
			continue
		}
		d := 0
		if n > 1 {
			if m == n*(n-1)/2 {
				d = 1
			} else {
				d = 2
			}
		}
		if d < k-1 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1581B(os.Stdin, os.Stdout) }
