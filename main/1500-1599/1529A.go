package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1529A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c, s := [101]int{}, 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			c[v]++
			s += v
		}
		ans := 0
		for i := 100; i*n > s; i-- {
			ans += c[i]
			n -= c[i]
			s -= i * c[i]
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1529A(os.Stdin, os.Stdout) }
