package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

func solve(reader io.Reader, writer io.Writer) {
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
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n, q, root := read(), read(), read()-1
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	vs := make([]int, 0, 2*n-1)
	pos := make([]int, n)
	depths := make([]int, 0, 2*n-1)
	var dfs func(v, fa, d int)
	dfs = func(v, fa, d int) {
		pos[v] = len(vs)
		vs = append(vs, v)
		depths = append(depths, d)
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v, d+1)
				vs = append(vs, v)
				depths = append(depths, d)
			}
		}
	}
	dfs(root, -1, 0)

	var st [][20]int
	stInit := func(a []int) {
		n := len(a)
		st = make([][20]int, n)
		for i := range st {
			st[i][0] = a[i]
		}
		for j := uint(1); 1<<j <= n; j++ {
			for i := 0; i+(1<<j)-1 < n; i++ {
				st[i][j] = min(st[i][j-1], st[i+(1<<(j-1))][j-1])
			}
		}
	}
	stQuery := func(l, r int) int {
		k := uint(bits.Len(uint(r-l+1)) - 1)
		return min(st[l][k], st[r-(1<<k)+1][k])
	}
	calcLCA := func(v, w int) int {
		pv, pw := pos[v], pos[w]
		if pv > pw {
			pv, pw = pw, pv
		}
		return vs[stQuery(pv, pw)]
	}

	stInit(depths)
	for ; q > 0; q-- {
		Fprintln(out, calcLCA(read()-1, read()-1)+1)
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
