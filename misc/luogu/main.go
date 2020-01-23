package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n1, n2, m int
	Fscan(in, &n1, &n2, &m)
	g := make([][]int, n1+n2)
	for i := 0; i < m; i++ {
		var v, w int
		Fscan(in, &v, &w)
		if v > n1 || w > n2 {
			continue
		}
		v--
		w--
		w += n1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	cnt := 0
	match := make([]int, n1+n2)
	for i := range match {
		match[i] = -1
	}
	var used []bool
	var f func(v int) bool
	f = func(v int) bool {
		for _, w := range g[v] {
			if !used[w] {
				used[w] = true
				if match[w] == -1 || f(match[w]) {
					match[w] = v
					return true
				}
			}
		}
		return false
	}
	for v := range g {
		used = make([]bool, n1+n2)
		if f(v) {
			cnt++
		}
	}
	Fprintln(out, cnt/2)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
