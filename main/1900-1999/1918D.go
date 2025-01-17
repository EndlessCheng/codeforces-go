package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1918D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		mx := 0
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			mx = max(mx, a[i])
		}
		ans := mx + sort.Search(1e9*n-mx, func(high int) bool {
			high += mx
			f := make([]int, n+2)
			q := []int{}
			s, l := 0, 0
			for i, x := range a {
				s += x
				for s > high {
					s -= a[l]
					l++
				}
				for len(q) > 0 && q[0] < l-1 {
					q = q[1:]
				}
				for len(q) > 0 && f[i]+x <= f[q[len(q)-1]]+a[q[len(q)-1]] {
					q = q[:len(q)-1]
				}
				q = append(q, i)
				f[i+1] = f[q[0]] + a[q[0]]
			}
			return f[n+1] <= high
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1918D(os.Stdin, os.Stdout) }
