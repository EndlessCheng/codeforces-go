package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf815C(in io.Reader, out io.Writer) {
	var n, b, v int
	Fscan(in, &n, &b)
	type pair struct{ nc, c int } // 不优惠，优惠
	a := make([]pair, n)
	g := make([][]int, n)
	Fscan(in, &a[0].nc, &a[0].c)
	for w := 1; w < n; w++ {
		Fscan(in, &a[w].nc, &a[w].c, &v)
		g[v-1] = append(g[v-1], w)
	}

	var dfs func(int) ([]pair, int)
	dfs = func(v int) ([]pair, int) {
		f := make([]pair, n+1)
		f[0].c = 1e18
		f[1] = pair{a[v].nc, a[v].nc - a[v].c}
		for i := 2; i <= n; i++ {
			f[i] = pair{1e18, 1e18}
		}
		size := 1
		for _, w := range g[v] {
			fw, sz := dfs(w)
			size += sz
			for j := size; j >= sz; j-- {
				for k, p := range fw {
					f[j].nc = min(f[j].nc, f[j-k].nc+p.nc)
					f[j].c = min(f[j].c, f[j-k].c+min(p.nc, p.c))
				}
			}
		}
		return f[:size+1], size
	}

	f, _ := dfs(0)
	for i := n; ; i-- {
		if min(f[i].nc, f[i].c) <= b {
			Fprint(out, i)
			return
		}
	}
}

//func main() { cf815C(bufio.NewReader(os.Stdin), os.Stdout) }
