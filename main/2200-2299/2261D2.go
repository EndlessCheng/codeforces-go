package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand/v2"
)

// https://github.com/EndlessCheng
func cf2261D2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		d := make([]int, n+2)
		mp := make([]int, m+1)
		for i := 1; i <= m; i++ {
			mp[i] = rand.Int()
		}

		for i := 1; i <= m; i++ {
			var l, r int
			Fscan(in, &l, &r)
			if l == 1 && r == n {
				continue
			}
			if r == n {
				r = l - 1
				l = 1
			}
			d[l] += mp[i]
			d[r+1] -= mp[i]
		}

		st := make([]int, n+1)
		pos := map[int]int{}
		for i := 1; i <= n; i++ {
			d[i] += d[i-1]
			if j, ok := pos[d[i]]; ok {
				st[j] = 1
				st[i] = 1
			}
			pos[d[i]] = i
		}
		st[0] = st[n]

		v := []int{}
		for i := 1; i <= n; i++ {
			if st[i] != 0 {
				v = append(v, i)
			}
		}

		g := []int{}
		for i := 0; i < len(v); i++ {
			if d[v[i]] == d[v[(i+1)%len(v)]] {
				g = append(g, v[i]+1)
			}
		}

		Fprintln(out, (len(g)+1)/2)
		for i := 0; i < (len(g)+1)/2; i++ {
			l := g[i]
			r := g[(i+len(g)/2)%len(g)] - 1
			if l <= r {
				Fprintln(out, l, r)
			} else {
				Fprintln(out, r+1, l-1)
			}
		}
	}
}

//func main() { cf2261D2(bufio.NewReader(os.Stdin), os.Stdout) }
