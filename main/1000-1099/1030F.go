package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type fenwick []int

func (f fenwick) update(i, v int) {
	for ; i < len(f); i += i & -i {
		f[i] += v
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func (f fenwick) kth(k int) (res int) {
	for b := 1 << 17; b > 0; b >>= 1 {
		if nxt := res | b; nxt < len(f) && f[nxt] < k {
			k -= f[nxt]
			res = nxt
		}
	}
	return res + 1
}

func cf1030F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007
	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] -= i
	}
	w := make([]int, n+1)
	fw := make(fenwick, n+1)
	fwa := make(fenwick, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &w[i])
		fw.update(i, w[i])
		fwa.update(i, w[i]*a[i]%mod)
	}

	for range q {
		Fscan(in, &l, &r)
		if l < 0 {
			l = -l
			d := r - w[l]
			w[l] = r
			fw.update(l, d)
			fwa.update(l, d*a[l]%mod)
		} else {
			m := fw.kth((fw.pre(l-1) + fw.pre(r) + 1) / 2)
			x := a[m]
			ans := fw.query(l, m)%mod*x - fwa.query(l, m) +
				   fwa.query(m, r) - fw.query(m, r)%mod*x
			Fprintln(out, (ans%mod+mod)%mod)
		}
	}
}

//func main() { cf1030F(bufio.NewReader(os.Stdin), os.Stdout) }
