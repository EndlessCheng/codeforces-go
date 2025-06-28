package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf627D(in io.Reader, out io.Writer) {
	var n, k, root int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < a[root] {
			root = i
		}
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := sort.Search(1e6, func(low int) bool {
		low++
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			full := true
			sum, mx, mx2 := 1, 0, 0
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				s := dfs(w, v)
				if s == 1e9 {
					return 1e9
				}
				if s > 0 {
					sum += s
				} else {
					full = false
					if -s > mx {
						mx2 = mx
						mx = -s
					} else {
						mx2 = max(mx2, -s)
					}
				}
			}
			if a[v] < low {
				return 0
			}
			sum += mx
			if sum+mx2 >= k {
				return 1e9
			}
			if full {
				return sum
			}
			return -sum
		}
		return dfs(root, -1) != 1e9
	})
	Fprint(out, ans)
}

//func main() { cf627D(bufio.NewReader(os.Stdin), os.Stdout) }
