package main

import (
	"bufio"
	. "fmt"
	"io"
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

	n, m := read(), read()
	g := make([][]int, n)
	rg := make([][]int, n)
	for i := 0; i < m; i++ {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v)
	}
	used := make([]bool, n)

	vs := make([]int, 0, n)
	var dfs func(int)
	dfs = func(v int) {
		used[v] = true
		for _, w := range g[v] {
			if !used[w] {
				dfs(w)
			}
		}
		vs = append(vs, v)
	}
	for v := range g {
		if !used[v] {
			dfs(v)
		}
	}

	components := [][]int{}
	used = make([]bool, n)
	var comp []int
	var rdfs func(int)
	rdfs = func(v int) {
		used[v] = true
		comp = append(comp, v)
		for _, w := range rg[v] {
			if !used[w] {
				rdfs(w)
			}
		}
	}
	for i := len(vs) - 1; i >= 0; i-- {
		if !used[vs[i]] {
			comp = []int{}
			rdfs(vs[i])
			components = append(components, comp)
		}
	}

	last := components[len(components)-1]
	used = make([]bool, n)
	rdfs(last[0])
	for _, use := range used {
		if !use {
			Fprint(out, 0)
			return
		}
	}
	Fprint(out, len(last))
}

func main() {
	solve(os.Stdin, os.Stdout)
}
