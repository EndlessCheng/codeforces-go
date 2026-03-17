package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1941E(in io.Reader, out io.Writer) {
	var T, n, m, k, d, v, f int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k, &d)
		rowRes := make([]int, n)
		ans, s := int(1e18), 0
		for i := range n {
			Fscan(in, &v)
			type pair struct{ j, f int }
			q := []pair{{0, 1}}
			for j := 1; j < m; j++ {
				Fscan(in, &v)
				for q[0].j < j-d-1 {
					q = q[1:]
				}
				f = q[0].f + v + 1
				for f <= q[len(q)-1].f {
					q = q[:len(q)-1]
				}
				q = append(q, pair{j, f})
			}
			rowRes[i] = f
			s += f
			if i >= k-1 {
				ans = min(ans, s)
				s -= rowRes[i-k+1]
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1941E(bufio.NewReader(os.Stdin), os.Stdout) }
