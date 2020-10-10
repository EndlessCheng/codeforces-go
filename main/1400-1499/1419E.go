package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1419E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ p, k int }

	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		op := 1
		fs := []pair{}
		for i := 2; i*i <= n; i++ {
			k := 0
			for ; n%i == 0; n /= i {
				k++
			}
			if k > 0 {
				fs = append(fs, pair{i, k})
				if k > 1 {
					op = 0
				}
			}
		}
		if n > 1 {
			fs = append(fs, pair{n, 1})
		}

		ans := []interface{}{}
		var f func(m, v int)
		f = func(m, v int) {
			if m == 0 {
				ans = append(ans, v)
				return
			}
			p := fs[bits.TrailingZeros(uint(m))]
			m &= m - 1
			for i := 0; i < p.k; i++ {
				v *= p.p
				f(m, v)
			}
		}
		n = len(fs)
		for i := 1; i < 1<<n; i++ {
			m := i ^ i>>1
			if m+1 == 1<<n {
				l = len(ans)
				f(m, 1)
				r = len(ans)
			} else {
				f(m, 1)
			}
		}
		if n == 2 {
			l++
		} else {
			op = 0
		}
		Fprint(out, ans[:l]...)
		Fprint(out, " ")
		Fprint(out, ans[r:]...)
		Fprint(out, " ")
		Fprintln(out, ans[l:r]...)
		Fprintln(out, op)
	}
}

//func main() { CF1419E(os.Stdin, os.Stdout) }
