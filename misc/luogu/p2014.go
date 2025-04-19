package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2014(in io.Reader, out io.Writer) {
	var n, m, p int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	g := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &p, &a[i])
		g[p] = append(g[p], i)
	}
	var dfs func(int) ([]int, int)
	dfs = func(v int) ([]int, int) {
		f := make([]int, n+2)
		size := 0
		for _, w := range g[v] {
			fw, sz := dfs(w)
			for j := size; j >= 0; j-- {
				for k, res := range fw {
					f[j+k+1] = max(f[j+k+1], f[j+1]+res)
				}
			}
			size += sz
		}
		size++
		for i := 1; i <= size; i++ {
			f[i] += a[v]
		}
		return f[:size+1], size
	}
	f, _ := dfs(0)
	Fprint(out, f[min(n, m)+1])
}

//func main() { p2014(bufio.NewReader(os.Stdin), os.Stdout) }
