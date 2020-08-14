package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1101D(_r io.Reader, _w io.Writer) {
	const mx int = 2e5
	divisors := [mx + 1][]int{}
	for i := 2; i <= mx; i++ {
		if len(divisors[i]) == 0 {
			for j := i; j <= mx; j += i {
				divisors[j] = append(divisors[j], i)
			}
		}
	}

	var n, v, w, ans int
	in := bufio.NewReader(_r)
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	// 关键：从根 v 往下看，根 v 上的各个质数 p 能「延伸」到多远
	var f func(v, fa int) map[int][2]int
	f = func(v, fa int) map[int][2]int {
		res := map[int][2]int{}
		for _, p := range divisors[a[v]] {
			res[p] = [2]int{}
		}
		for _, w := range g[v] {
			if w != fa {
				sub := f(w, v)
				for p, dis := range res {
					d := sub[p][0]
					if d <= dis[1] {
						continue
					}
					if d > dis[0] {
						dis[0], d = d, dis[0]
					}
					dis[1] = d
					res[p] = dis
				}
			}
		}
		for p, dis := range res {
			if v := dis[0] + dis[1] + 1; v > ans {
				ans = v
			}
			dis[0]++
			res[p] = dis
		}
		return res
	}
	f(0, -1)
	Fprint(_w, ans)
}

//func main() { CF1101D(os.Stdin, os.Stdout) }
