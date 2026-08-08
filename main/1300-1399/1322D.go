package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1322D(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	l := make([]int, n+1)
	s := make([]int, n+1)
	c := make([]int, n+m+2)
	for i := n; i >= 1; i-- {
		Fscan(in, &l[i])
	}
	for i := n; i >= 1; i-- {
		Fscan(in, &s[i])
	}
	for i := 1; i <= n+m; i++ {
		Fscan(in, &c[i])
	}

	f := make([][]int, n+m+2)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = -1e18
		}
		f[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		t := l[i]
		for j := n; j >= 1; j-- {
			f[t][j] = max(f[t][j], f[t][j-1]-s[i]+c[t])
		}
		k := n
		for j := t; j <= n+m; j++ {
			for x := 0; x <= k; x++ {
				f[j+1][x>>1] = max(
					f[j+1][x>>1],
					f[j][x]+c[j+1]*(x>>1),
				)
			}
			k >>= 1
		}
	}
	Fprint(out, f[n+m][0])
}

//func main() { cf1322D(bufio.NewReader(os.Stdin), os.Stdout) }
