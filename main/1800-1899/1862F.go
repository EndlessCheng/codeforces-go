package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1862F(in io.Reader, out io.Writer) {
	var T, n, v, p, q int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &p, &q, &n)
		mx := n * 1e4
		f := make([]bool, mx+1)
		f[0] = true
		s := 0
		for range n {
			Fscan(in, &v)
			s += v
			for j := s; j >= v; j-- {
				f[j] = f[j] || f[j-v]
			}
		}

		ans := mx * 2
		for i, ok := range f {
			if ok {
				ans = min(ans, max((i+p-1)/p, (s-i+q-1)/q))
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1862F(bufio.NewReader(os.Stdin), os.Stdout) }
