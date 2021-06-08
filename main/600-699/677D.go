package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF677D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	type pair struct{ x, y, d int }

	var n, m, p, v int
	Fscan(in, &n, &m, &p)
	pos := make([][]pair, n*m+1)
	pos[0] = []pair{{}}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fscan(in, &v)
			pos[v] = append(pos[v], pair{i, j, 1e9})
		}
	}
	for i, ps := range pos[:p] {
		sort.Slice(ps, func(i, j int) bool { return ps[i].d < ps[j].d })
		if len(ps) > n+m {
			// 也可以全局 BFS
			ps = ps[:n+m]
		}
		for _, p := range ps {
			for k, q := range pos[i+1] {
				if d := p.d + abs(p.x-q.x) + abs(p.y-q.y); d < q.d {
					pos[i+1][k].d = d
				}
			}
		}
	}
	Fprint(out, pos[p][0].d)
}

//func main() { CF677D(os.Stdin, os.Stdout) }
