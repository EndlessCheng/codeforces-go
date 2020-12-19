package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1463E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v, w int
	Fscan(in, &n, &k)
	fa := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		fa[x] = y
		return true
	}
	same := func(x, y int) bool { return find(x) == find(y) }
	p := make([]int, n)
	to := make([]int, n)
	for i := range p {
		Fscan(in, &p[i])
		p[i]--
		to[i] = -1
		fa[i] = i
	}
	for ; k > 0; k-- {
		Fscan(in, &v, &w)
		v--
		w--
		if !merge(w, v) {
			Fprint(out, 0)
			return
		}
		to[v] = w
	}

	// 将链缩为点
	g := make([][]int, n)
	deg := make([]int, n)
	for w, v := range p {
		if v >= 0 && !same(v, w) {
			v, w = fa[v], fa[w]
			g[v] = append(g[v], w)
			deg[w]++
		}
	}

	// 跑缩点后的拓扑排序
	q := []int{}
	for i, d := range deg {
		if d == 0 && fa[i] == i {
			q = append(q, i)
		}
	}
	orders := []int{}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for x := v; x >= 0; x = to[x] {
			orders = append(orders, x)
		}
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}

	if len(orders) < n {
		Fprint(out, 0)
		return
	}
	pos := make([]int, n)
	for i, o := range orders {
		pos[o] = i
	}
	// 由于把链缩为了点，还需要额外判断链上的拓扑序是否和树上的父子顺序一致
	// 这一步也可以放在建图之前，见 https://codeforces.com/contest/1463/submission/101602043
	for w, v := range p {
		if v >= 0 && pos[v] > pos[w] {
			Fprint(out, 0)
			return
		}
	}
	for _, v := range orders {
		Fprint(out, v+1, " ")
	}
}

func CF1463ESolution2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v, w, rt int
	Fscan(in, &n, &k)
	p := make([]int, n)
	to := make([]int, n)
	top := make([]int, n)
	for i := range p {
		Fscan(in, &p[i])
		if p[i]--; p[i] < 0 {
			rt = i
		}
		to[i] = -1
		top[i] = i
	}
	inChain := make([]bool, n)
	for ; k > 0; k-- {
		Fscan(in, &v, &w)
		v--
		w--
		to[v] = w
		inChain[w] = true
	}

	g := make([][]int, n)
	deg := make([]int, n)
	for i, c := range inChain {
		if !c {
			for v := i; v >= 0; v = to[v] {
				top[v] = i
				if p[v] >= 0 && top[p[v]] != i {
					g[p[v]] = append(g[p[v]], i)
					deg[i]++
				}
			}
		}
	}

	ans := []interface{}{}
	q := []int{}
	// rt 入度不为零的数据：
	// 3 1
	// 0 3 1
	// 1 2
	if !inChain[rt] && deg[rt] == 0 {
		q = []int{rt}
	}
	for len(q) > 0 {
		for v, q = q[0], q[1:]; v >= 0; v = to[v] {
			ans = append(ans, v+1)
			for _, w := range g[v] {
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
	}
	if len(ans) < n {
		Fprint(out, 0)
	} else {
		Fprint(out, ans...)
	}
}

//func main() { CF1463E(os.Stdin, os.Stdout) }
