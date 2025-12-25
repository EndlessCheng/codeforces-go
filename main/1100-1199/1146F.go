package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1146F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, p int
	Fscan(in, &n)
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &p)
		g[p-1] = append(g[p-1], w)
	}

	// 0: 子树中没有路径延伸到 v
	// 1: 子树中恰好有一条路径延伸到 v
	// 2: 子树中有 >= 两条路径在 v 处合并，或者 v 本身就是一个叶子
	f := make([][3]int, n)
	for v := n - 1; v >= 0; v-- {
		if g[v] == nil {
			f[v][2] = 1
			continue
		}
		f[v][0] = 1
		for _, w := range g[v] {
			zero := f[w][0] + f[w][2]
			one := f[w][1] + f[w][2]
			f[v][2] = (f[v][2]*(zero+one) + f[v][1]*one) % mod
			f[v][1] = (f[v][1]*zero + f[v][0]*one) % mod
			f[v][0] = f[v][0] * zero % mod
		}
	}
	Fprint(out, (f[0][0]+f[0][2])%mod)
}

//func main() { cf1146F(bufio.NewReader(os.Stdin), os.Stdout) }
