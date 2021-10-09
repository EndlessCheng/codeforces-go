package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1566E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := 1
		var f func(int, int) bool
		f = func(v, fa int) bool {
			c := 0
			for _, w := range g[v] {
				if w != fa && f(w, v) {
					c++
				}
			}
			if c == 0 { // 叶子，或者删掉所有子 bud 后为叶子
				return true
			}
			ans += c - 1 // -1 是因为 bud 接在某个叶子下
			return false // 删掉该 bud
		}
		f(1, 0)
		Fprintln(out, ans)
	}
}

//func main() { CF1566E(os.Stdin, os.Stdout) }
