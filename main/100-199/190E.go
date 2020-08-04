package _00_199

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol190E(reader io.Reader, writer io.Writer) {
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
	var fa []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }
	same := func(x, y int) bool { return find(x) == find(y) }

	n, m := read(), read()
	g := make([][]int, n)
	for ; m > 0; m-- {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	maxDeg0, maxDeg0V := 0, 0
	for v, edges := range g {
		if deg0 := n - 1 - len(edges); deg0 > maxDeg0 {
			maxDeg0 = deg0
			maxDeg0V = v
		}
	}
	if maxDeg0 == 0 {
		Fprintln(out, n)
		for i := 1; i <= n; i++ {
			Fprintln(out, 1, i)
		}
		return
	}

	merge0 := func(v int, edges []int) {
		vs := map[int]bool{v: true}
		for _, w := range edges {
			vs[w] = true
		}
		for i := 0; i < n; i++ {
			if !vs[i] {
				merge(i, v)
			}
		}
	}
	initFa(n)
	merge0(maxDeg0V, g[maxDeg0V])
	for v, edges := range g {
		if !same(v, maxDeg0V) {
			merge0(v, edges)
		}
	}

	cnt := map[int][]interface{}{}
	for i := range fa {
		v := find(i)
		cnt[v] = append(cnt[v], i+1)
	}
	Fprintln(out, len(cnt))
	for _, vs := range cnt {
		Fprint(out, len(vs), " ")
		Fprintln(out, vs...)
	}
}

//func main() {
//	Sol190E(os.Stdin, os.Stdout)
//}
