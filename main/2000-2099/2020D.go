package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2020D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, a, d, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		fa := [11][]int{}
		for i := range fa {
			fa[i] = make([]int, n+i+1)
			for j := range fa[i] {
				fa[i][j] = j
			}
		}
		find := func(d, x int) int {
			fa := fa[d]
			rt := x
			for fa[rt] != rt {
				rt = fa[rt]
			}
			for fa[x] != rt {
				fa[x], x = rt, fa[x]
			}
			return rt
		}

		for range m {
			Fscan(in, &a, &d, &k)
			to := find(d, a+k*d)
			for i := find(d, a); i <= a+k*d; i = find(d, i+d) {
				fa[d][find(d, i)] = to
			}
		}

		cc := n
		for i := 1; i < len(fa); i++ {
			for j := 1; j <= n; j++ {
				x := find(0, j)
				y := find(0, find(i, j))
				if x != y {
					fa[0][x] = y
					cc--
				}
			}
		}
		Fprintln(out, cc)
	}
}

//func main() { cf2020D(bufio.NewReader(os.Stdin), os.Stdout) }
