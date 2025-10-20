package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf868F(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	cnt := make([]int, n+1)
	s, L, R := 0, 0, -1
	upd := func(ql, qr int) {
		for ql < L {
			L--
			s += cnt[a[L]]
			cnt[a[L]]++
		}
		for qr > R {
			R++
			s += cnt[a[R]]
			cnt[a[R]]++
		}
		for ; ql > L; L++ {
			cnt[a[L]]--
			s -= cnt[a[L]]
		}
		for ; qr < R; R-- {
			cnt[a[R]]--
			s -= cnt[a[R]]
		}
	}

	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = 1e18
		}
	}
	var dfs func(int, int, int, int, int)
	dfs = func(i, ql, qr, optL, optR int) {
		if ql > qr {
			return
		}
		mid := (ql + qr) / 2
		mn := int(1e18)
		opt := -1
		for j := min(mid, optR); j >= optL; j-- {
			upd(j+1, mid)
			if s+f[i-1][j] < mn {
				mn = s + f[i-1][j]
				opt = j
			}
		}
		f[i][mid] = mn
		dfs(i, ql, mid-1, optL, opt)
		dfs(i, mid+1, qr, opt, optR)
	}

	f[0][0] = 0
	for i := 1; i <= k; i++ {
		dfs(i, 0, n, 0, n)
	}
	Fprint(out, f[k][n])
}

//func main() { cf868F(bufio.NewReader(os.Stdin), os.Stdout) }
