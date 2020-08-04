package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func Sol832D(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}
	maxs := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val > ans {
				ans = val
			}
		}
		return ans
	}

	n, q := read(), read()
	g := make([][]int, n)
	for v := 1; v < n; v++ {
		w := read() - 1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vs := make([]int, 0, 2*n-1)
	pos := make([]int, n)
	depths := make([]int, 0, 2*n-1)
	dis := make([]int, n)
	var dfs func(v, fa, d int)
	dfs = func(v, fa, d int) {
		pos[v] = len(vs)
		vs = append(vs, v)
		depths = append(depths, d)
		dis[v] = d
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
				vs = append(vs, v)
				depths = append(depths, d)
			}
		}
	}
	dfs(0, -1, 0)

	type pair struct{ v, i int }
	var st [][18]pair
	stInit := func(a []int) {
		n := len(a)
		st = make([][18]pair, n)
		for i := range st {
			st[i][0] = pair{a[i], i}
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st0, st1 := st[i][j-1], st[i+(1<<(j-1))][j-1]
				if st0.v < st1.v {
					st[i][j] = st0
				} else {
					st[i][j] = st1
				}
			}
		}
	}
	stQuery := func(l, r int) int {
		k := uint(bits.Len(uint(r-l+1)) - 1)
		st0, st1 := st[l][k], st[r-(1<<k)+1][k]
		if st0.v < st1.v {
			return st0.i
		}
		return st1.i
	}
	lca := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return vs[stQuery(pv, pw)]
	}
	calcDis := func(v, w int) int {
		return dis[v] + dis[w] - 2*dis[lca(v, w)]
	}

	stInit(depths)
	for ; q > 0; q-- {
		a, b, c := read()-1, read()-1, read()-1
		v := lca(a, b) ^ lca(a, c) ^ lca(b, c)
		Fprintln(out, 1+maxs(calcDis(v, a), calcDis(v, b), calcDis(v, c)))
	}
}

//func main() {
//	Sol832D(os.Stdin, os.Stdout)
//}
