package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1242C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n)
	a := make([][]int64, n)
	id := map[int64]int{}
	sum := make([]int64, n)
	tot := int64(0)
	for i := range a {
		Fscan(in, &m)
		a[i] = make([]int64, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			id[a[i][j]] = i + 1
			sum[i] += a[i][j]
		}
		tot += sum[i]
	}
	if tot%int64(n) != 0 {
		Fprint(out, "No")
		return
	}
	tot /= int64(n)

	pick := make([]int64, n)
	to := make([]int, n)
	vis := make([]bool, 1<<n)
	var f func(int, int) bool
	f = func(s, p int) bool {
		if p == n {
			return true
		}
		if vis[s] {
			return false
		}
		if s>>p&1 > 0 {
			return f(s, p+1)
		}
		s |= 1 << p
	o:
		for _, v := range a[p] {
			pick[p] = v
			t, q, cur := s, p, v
			for {
				w := tot - sum[q] + cur
				if w == v {
					to[p] = q + 1
					break
				}
				q = id[w] - 1
				if q < 0 || t>>q&1 > 0 {
					continue o
				}
				pick[q], to[q] = w, id[cur]
				t |= 1 << q
				cur = w
			}
			if f(t, p+1) {
				return true
			}
			vis[t] = true
		}
		return false
	}
	if f(0, 0) {
		Fprintln(out, "Yes")
		for i, v := range pick {
			Fprintln(out, v, to[i])
		}
	} else {
		Fprint(out, "No")
	}
}

//func main() { CF1242C(os.Stdin, os.Stdout) }
