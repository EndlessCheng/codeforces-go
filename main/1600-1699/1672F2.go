package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1672F2(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}

		g := make([][]int, n+1)
		cnt := make([]int, n+1)
		mx := 1
		for i := 1; i <= n; i++ {
			var x int
			Fscan(in, &x)
			g[a[i]] = append(g[a[i]], x)
			cnt[a[i]]++
			if cnt[a[i]] > cnt[mx] {
				mx = a[i]
			}
		}

		color := make([]int8, n+1)
		color[mx] = 2

		var dfs func(int) bool
		dfs = func(x int) bool {
			color[x] = 1
			for _, y := range g[x] {
				if color[y] == 1 || color[y] == 0 && !dfs(y) {
					return false
				}
			}
			color[x] = 2
			return true
		}

		for i := 1; i <= n; i++ {
			if color[i] == 0 && !dfs(i) {
				Fprintln(out, "WA")
				continue o
			}
		}
		Fprintln(out, "AC")
	}
}

//func main() { cf1672F2(bufio.NewReader(os.Stdin), os.Stdout) }
