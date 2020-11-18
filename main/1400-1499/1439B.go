package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1439B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, k, v, w int
t:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		g := make([]map[int]bool, n)
		for i := range g {
			g[i] = map[int]bool{}
		}
		deg := make([]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			v--
			w--
			g[v][w] = true
			g[w][v] = true
			deg[v]++
			deg[w]++
		}

		left := n
		del := make([]bool, n)
		q := []int{}
		for i, d := range deg {
			if d < k {
				del[i] = true
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v, q = q[0], q[1:]
			left--
			if deg[v] == k-1 {
				// 检查 v 和它的小伙伴们(邻居)能否组成 k-团
				clique := []int{}
				for w := range g[v] {
					if deg[w] >= k-1 {
						clique = append(clique, w)
					}
				}
				if len(clique) == k-1 {
					for i, v := range clique {
						for _, w := range clique[:i] {
							if !g[v][w] {
								goto o
							}
						}
					}
					Fprint(out, "2\n", v+1, " ")
					for _, v := range clique {
						Fprint(out, v+1, " ")
					}
					Fprintln(out)
					continue t
				}
			}
		o:
			deg[v] = 0
			for w := range g[v] {
				if deg[w]--; deg[w] < k && !del[w] {
					del[w] = true
					q = append(q, w)
				}
			}
		}

		if left > 0 {
			Fprintln(out, 1, left)
			for i, d := range del {
				if !d {
					Fprint(out, i+1, " ")
				}
			}
			Fprintln(out)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { CF1439B(os.Stdin, os.Stdout) }
