package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1282E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pr struct{ x, y int }
	sort3 := func(a ...int) (x, y, z int) { sort.Ints(a); return a[0], a[1], a[2] }

	var T, n, x, y, z int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		id := make([]map[pr]int, n)
		for i := range id {
			id[i] = map[pr]int{}
		}
		cnt := make([]int, n)
		for i := 1; i < n-1; i++ {
			Fscan(in, &x, &y, &z)
			x, y, z = sort3(x-1, y-1, z-1)
			id[x][pr{y, z}] = i
			id[y][pr{x, z}] = i
			id[z][pr{x, y}] = i
			cnt[x]++
			cnt[y]++
			cnt[z]++
		}
		if n == 3 {
			Fprintln(out, 1, 2, 3)
			Fprintln(out, 1)
			continue
		}

		ans := make([]int, 0, n-2)
		g := make([][]int, n)
		inner := map[pr]bool{}
		add := func(v, w int) {
			if v > w {
				v, w = w, v
			}
			if !inner[pr{v, w}] {
				g[v] = append(g[v], w)
				g[w] = append(g[w], v)
			}
		}
		q := []int{}
		for i, c := range cnt {
			if c == 1 {
				q = append(q, i)
			}
		}
		for n -= 2; n > 0; n-- {
			x := q[0]
			q = q[1:]
			p, i := pr{}, 0
			for p, i = range id[x] {
				break
			}
			ans = append(ans, i)

			y, z := p.x, p.y
			if cnt[y]--; cnt[y] == 1 {
				q = append(q, y)
			}
			if cnt[z]--; cnt[z] == 1 {
				q = append(q, z)
			}

			add(x, y)
			add(x, z)
			inner[p] = true

			x, y, z = sort3(x, y, z)
			delete(id[x], pr{y, z})
			delete(id[y], pr{x, z})
			delete(id[z], pr{x, y})
		}
		fa, v := -1, 0
		for {
			Fprint(out, v+1, " ")
			for _, w := range g[v] {
				if w != fa {
					fa, v = v, w
					break
				}
			}
			if v == 0 {
				break
			}
		}
		Fprintln(out)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1282E(os.Stdin, os.Stdout) }
