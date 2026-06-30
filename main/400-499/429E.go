package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf429E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	type pair struct{ x, id int }
	ps := make([]pair, 0, 2*n)
	memo := make([]int, n)
	for i := range n {
		var a, b int
		Fscan(in, &a, &b)
		ps = append(ps, pair{a * 2, i * 2}, pair{b*2 + 1, i*2 + 1})
		memo[i] = -1
	}
	slices.SortFunc(ps, func(a, b pair) int { return a.x - b.x })

	p := make([]int, 2*n)
	for i := range n {
		a := ps[i*2].id
		b := ps[i*2+1].id
		p[a] = b
		p[b] = a
	}

	var dfs func(int, int) int
	dfs = func(i, c int) int {
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = c
		dfs(p[2*i]/2, c^(p[2*i]&1)^1)
		dfs(p[2*i+1]/2, c^(p[2*i+1]&1))
		return c
	}

	for i := range n {
		Fprint(out, dfs(i, 0), " ")
	}
}

//func main() { cf429E(bufio.NewReader(os.Stdin), os.Stdout) }
