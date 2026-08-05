package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2233D(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pos := map[int][]int{}
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}

		check := func(p, q int) bool {
			if p < 0 || q == n {
				return false
			}
			a[p], a[q] = a[q], a[p]
			vis := map[int]bool{}
			for i, v := range a {
				if i == 0 || v != a[i-1] {
					if vis[v] {
						a[p], a[q] = a[q], a[p]
						return false
					}
					vis[v] = true
				}
			}
			return true
		}

		for _, ps := range pos {
			m := len(ps)
			l, r := ps[0], ps[m-1]
			if r-l+1 == m {
				continue
			}
			empty := []int{}
			for i := 1; i < m; i++ {
				if ps[i]-ps[i-1] > 1 {
					empty = append(empty, ps[i-1]+1, ps[i]-1)
				}
			}
			if check(l-1, r) || check(l, r+1) || check(empty[0], r) || check(l, empty[len(empty)-1]) {
				break
			}
			Fprintln(out, "NO")
			continue o
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf2233D(bufio.NewReader(os.Stdin), os.Stdout) }
