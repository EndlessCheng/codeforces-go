package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	f := []int{}
	var dfs func(int, int)
	dfs = func(v, fa int) {
		x := a[v]
		j := sort.SearchInts(f, x)
		if j < len(f) {
			old := f[j]
			f[j] = x
			defer func() { f[j] = old }()
		} else {
			f = append(f, x)
			defer func() { f = f[:len(f)-1] }()
		}
		ans[v] = len(f)
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
			}
		}
	}
	dfs(0, -1)
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
