package main

import (
	. "fmt"
	"io"
	"slices"
)

func cf46E(in io.Reader, out io.Writer) {
	var n, m, v int
	Fscan(in, &n, &m)
	f := make([]int, m)
	for i := 0; i < n; i++ {
		if i%2 > 0 {
			for j := m - 2; j > 0; j-- {
				f[j] = max(f[j], f[j+1])
			}
			s := 0
			for j := 0; j < m-1; j++ {
				Fscan(in, &v)
				s += v
				f[j] = f[j+1] + s
			}
			Fscan(in, &v)
			f[m-1] = -1e18
		} else {
			Fscan(in, &v)
			s := v
			mx := f[0]
			f[0] = -1e18
			for j := 1; j < m; j++ {
				Fscan(in, &v)
				s += v
				f[j], mx = mx+s, max(mx, f[j])
			}
		}
	}
	Fprint(out, slices.Max(f))
}

//func main() { cf46E(bufio.NewReader(os.Stdin), os.Stdout) }
