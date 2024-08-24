package main

import (
	. "fmt"
	"io"
)

func cf1650G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var T, n, m, s, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s, &t)
		s--
		t--
		g := make([][]int, n)
		for ; m > 0; m-- {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		// 偶数下标表示最短路，奇数下标表示次短路
		f := make([]int, n*2)
		dis := make([]int, n*2)
		for i := range dis {
			dis[i] = -1
		}
		dis[s*2] = 0
		f[s*2] = 1
		q := []int{s * 2}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			newD := dis[v] + 1
			for _, w := range g[v/2] {
				w *= 2
				if dis[w] < 0 {
					dis[w] = newD
					f[w] = f[v]
					q = append(q, w)
				} else if newD == dis[w] {
					f[w] = (f[w] + f[v]) % mod
				} else if newD == dis[w]+1 { // 次短路
					w++
					if dis[w] < 0 {
						dis[w] = newD
						q = append(q, w)
					}
					f[w] = (f[w] + f[v]) % mod
				}
			}
		}
		Fprintln(out, (f[t*2]+f[t*2+1])%mod)
	}
}

//func main() { cf1650G(bufio.NewReader(os.Stdin), os.Stdout) }
