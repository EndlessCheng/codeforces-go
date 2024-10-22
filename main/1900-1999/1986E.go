package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1986E(in io.Reader, out io.Writer) {
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		g := map[int][]int{}
		for range n {
			Fscan(in, &v)
			g[v%k] = append(g[v%k], v/k)
		}

		ans := 0
		odd := false
		for _, a := range g {
			slices.Sort(a)
			m := len(a)
			suf := 0
			for i := m - 2; i >= 0; i -= 2 {
				suf += a[i+1] - a[i]
			}
			if m%2 == 0 {
				ans += suf
				continue
			}

			if odd {
				ans = -1
				break
			}
			odd = true

			mn := suf
			pre := 0
			for i := 1; i < m; i += 2 {
				pre += a[i] - a[i-1]
				suf -= a[i+1] - a[i]
				mn = min(mn, pre+suf)
			}
			ans += mn
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1986E(bufio.NewReader(os.Stdin), os.Stdout) }
