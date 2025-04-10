package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func p3957(in io.Reader, out io.Writer) {
	var n, d, k int
	Fscan(in, &n, &d, &k)
	type pair struct{ x, pt int }
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].x, &a[i].pt)
	}
	lim := a[n].x - d
	f := make([]int, n+1)
	ans := sort.Search(lim+1, func(g int) bool {
		l := max(d-g, 1)
		q := []int{0}
		j := 1
		for i := 1; i <= n; i++ {
			if a[i].x < l {
				f[i] = -1e18
				continue
			}
			for a[j].x <= a[i].x-l {
				for len(q) > 0 && f[j] >= f[q[len(q)-1]] {
					q = q[:len(q)-1]
				}
				q = append(q, j)
				j++
			}
			for len(q) > 0 && a[q[0]].x < a[i].x-d-g {
				q = q[1:]
			}
			if len(q) > 0 {
				f[i] = f[q[0]] + a[i].pt
				if f[i] >= k {
					return true
				}
			} else {
				f[i] = -1e18
			}
		}
		return false
	})
	if ans > lim {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { p3957(bufio.NewReader(os.Stdin), os.Stdout) }
