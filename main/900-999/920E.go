package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func Sol920E(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
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

	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
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
		Fprintln(out, strings.Repeat("1 ", n))
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

	cnt := map[int]int{}
	for i := range fa {
		cnt[find(i)]++
	}
	values := make([]int, 0, len(cnt))
	for _, v := range cnt {
		values = append(values, v)
	}
	sort.Ints(values)
	Fprintln(out, len(values))
	for _, v := range values {
		Fprint(out, v, " ")
	}
}

//func main() {
//	Sol920E(os.Stdin, os.Stdout)
//}
