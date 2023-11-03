package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1634F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, mod, v, cnt0, l, r int
	var op string
	Fscan(in, &n, &q, &mod)
	f := make([]int, n+2)
	a := make([]int, n+1)
	Fscan(in, &a[1])
	f[1] = 1
	for i := 2; i <= n; i++ {
		Fscan(in, &a[i])
		f[i] = (f[i-1] + f[i-2]) % mod
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		a[i] = (a[i] - v) % mod
	}
	for i := n; i > 1; i-- {
		a[i] = (a[i] - a[i-1] - a[i-2]) % mod
		if a[i] == 0 {
			cnt0++
		}
	}
	if a[1] == 0 {
		cnt0++
	}

	update := func(i, add int) {
		if i > n {
			return
		}
		if a[i] == 0 {
			cnt0--
		}
		a[i] = (a[i] + add) % mod
		if a[i] == 0 {
			cnt0++
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		sgn := int(op[0]&1*2) - 1
		update(l, sgn)
		update(r+1, -f[r-l+2]*sgn)
		update(r+2, -f[r-l+1]*sgn)
		if cnt0 == n {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1634F(os.Stdin, os.Stdout) }
