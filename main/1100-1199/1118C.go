package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1118C(in io.Reader, out io.Writer) {
	var n, v, p int
	Fscan(in, &n)
	cnt := map[int]int{}
	for i := 0; i < n*n; i++ {
		Fscan(in, &v)
		cnt[v]++
	}
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	m := n / 2
	add := func(i, j, v int) bool {
		if ans[i][j] > 0 || ans[i][n-1-j] > 0 || ans[n-1-i][j] > 0 || ans[n-1-i][n-1-j] > 0 {
			return false
		}
		ans[i][j] = v
		ans[i][n-1-j] = v
		ans[n-1-i][j] = v
		ans[n-1-i][n-1-j] = v
		return true
	}
	if n&1 > 0 {
		q := 0
		addMid := func(v int) bool {
			return q < m && add(q, m, v) || q < 2*m && add(m, q-m, v)
		}
		for v, c := range cnt {
			if c&3 == 1 || c&3 == 3 {
				if !add(m, m, v) {
					Fprint(out, "NO")
					return
				}
				if c&3 == 3 {
					if !addMid(v) {
						Fprint(out, "NO")
						return
					}
					q++
				}
			} else if c&3 == 2 {
				if !addMid(v) {
					Fprint(out, "NO")
					return
				}
				q++
			}
			cnt[v] -= c & 3
		}
		for v, c := range cnt {
			if c&3 == 0 {
				for ; c > 0 && q < 2*m; c -= 2 {
					addMid(v)
					q++
				}
				for ; c > 0; c -= 4 {
					add(p/m, p%m, v)
					p++
				}
			}
		}
	} else {
		for v, c := range cnt {
			if c&3 > 0 {
				Fprint(out, "NO")
				return
			}
			for ; c > 0; c -= 4 {
				add(p/m, p%m, v)
				p++
			}
		}
	}
	Fprintln(out, "YES")
	for _, r := range ans {
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1118C(os.Stdin, os.Stdout) }
