package main

import (
	. "fmt"
	"io"
)

func cf1997D(in io.Reader, out io.Writer) {
	var T, n, p int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		for w := 1; w < n; w++ {
			Fscan(in, &p)
			g[p-1] = append(g[p-1], w)
		}
		var dfs func(int) int
		dfs = func(v int) int {
			if g[v] == nil {
				return a[v]
			}
			mn := int(1e9)
			for _, w := range g[v] {
				mn = min(mn, dfs(w))
			}
			if a[v] > mn {
				return mn
			}
			return (a[v] + mn) / 2
		}
		mn := int(1e9)
		for _, w := range g[0] {
			mn = min(mn, dfs(w))
		}
		Fprintln(out, a[0]+mn)
	}
}

//func main() { cf1997D(bufio.NewReader(os.Stdin), os.Stdout) }
