package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF56E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, posR int
	Fscan(in, &n)
	fa := make([]int, n)
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	a := make([]struct{ x, h, i int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].h)
		a[i].i = i
		fa[i] = i
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	type pair struct{ v, i int }
	stk := []pair{{2e9, n}}
	for i := n - 1; i >= 0; i-- {
		v := a[i].x + a[i].h
		for {
			if top := stk[len(stk)-1]; top.v > v {
				posR = top.i
				break
			}
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, pair{v, i})
		j := sort.Search(n, func(i int) bool { return a[i].x >= v }) - 1
		if j >= posR {
			j = posR
		}
		fa[f(i)] = f(j)
	}
	ans := make([]interface{}, n)
	for i, p := range a {
		ans[p.i] = f(i) - i + 1
	}
	Fprintln(out, ans...)
}

//func main() { CF56E(os.Stdin, os.Stdout) }
