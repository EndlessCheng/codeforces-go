package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1400B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var T, c, d, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &c, &d, &n, &m, &v, &w)
		if v > w {
			n, m, v, w = m, n, w, v
		}
		ans := 0
		for i := 0; i <= n && i*v <= c; i++ {
			j := min(n-i, d/v)
			if s := i + j + min(m, (c-i*v)/w+(d-j*v)/w); s > ans {
				ans = s
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1400B(os.Stdin, os.Stdout) }
