package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1654C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int64
	for Fscan(in, &T); T > 0; T-- {
		c := map[int64]int{}
		s := int64(0)
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &v)
			c[v]++
			s += v
		}
		var f func(int64) bool
		f = func(x int64) bool {
			if c[x] > 0 {
				c[x]--
				return true
			}
			return x > 1 && f(x/2) && f((x+1)/2)
		}
		if f(s) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1654C(os.Stdin, os.Stdout) }
