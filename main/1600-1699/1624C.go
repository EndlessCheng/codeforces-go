package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1624C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v uint
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		l := bits.Len(n)
		c := make([]int, n+1)
		for i := uint(0); i < n; i++ {
			Fscan(in, &v)
			if d := bits.Len(v) - l; d > 0 {
				v >>= d
			}
			if v > n {
				v >>= 1
			}
			c[v]++
		}
		for ; n > 0; n-- {
			if c[n] == 0 {
				Fprintln(out, "NO")
				continue o
			}
			c[n>>1] += c[n] - 1
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1624C(os.Stdin, os.Stdout) }
