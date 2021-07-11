package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF258C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int64 = 1e9 + 7
	pow := func(x, n int) (res int64) {
		res = 1
		for x := int64(x); n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}

	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)

	mx := a[n-1]
	p := make([]int, mx+1)
	cur := 1
	for i, v := range a {
		for ; cur <= v; cur++ {
			p[cur] = i
		}
	}
	ds := make([][]int, mx+1)
	for i := 1; i <= mx; i++ {
		for j := i; j <= mx; j += i {
			ds[j] = append(ds[j], i)
		}
	}

	ans := int64(1)
	for i := 2; i <= mx; i++ {
		mul := int64(1)
		d := ds[i]
		nd := len(d)
		for j := 1; j < nd-1; j++ {
			mul = mul * pow(j, p[d[j]]-p[d[j-1]]) % mod
		}
		p, q := p[d[nd-2]], p[d[nd-1]]
		m1 := mul * pow(nd-1, n-p) % mod
		mul = mul * pow(nd-1, q-p) % mod
		m2 := mul * pow(nd, n-q) % mod
		ans += m2 - m1
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF258C(os.Stdin, os.Stdout) }
