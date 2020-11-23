package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF559C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	const mx int = 2e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	pow := func(x, n int64) (res int64) {
		res = 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 { return F[n] * invF[k] % mod * invF[n-k] % mod }

	var h, w, n int
	Fscan(in, &h, &w, &n)
	type pair struct{ x, y int }
	a := make([]pair, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i].x, &a[i].y)
	}
	a[n] = pair{h, w}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
	f := make([]int64, n+1)
	for i, p := range a {
		f[i] = C(p.x+p.y-2, p.x-1)
		for j, q := range a[:i] {
			if q.x <= p.x && q.y <= p.y {
				f[i] -= f[j] * C(p.x-q.x+p.y-q.y, p.x-q.x) % mod
			}
		}
		f[i] = (f[i]%mod + mod) % mod
	}
	Fprint(out, f[n])
}

//func main() { CF559C(os.Stdin, os.Stdout) }
