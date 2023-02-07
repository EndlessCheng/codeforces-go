package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1361B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	var T, n, p int
	pow := func(n int) (res int64) {
		res = 1
		for x := int64(p); n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return
	}
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &p)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if p == 1 {
			Fprintln(out, n%2)
			continue
		}
		sort.Ints(a)

		targetK := a[n-1]
		type pair struct{ k, c int }
		s := []pair{}
		for i := n - 2; i >= 0; i-- {
			k := a[i]
			for len(s) > 0 && s[len(s)-1] == (pair{k, p - 1}) {
				s = s[:len(s)-1]
				k++
			}
			if k == targetK {
				if i == 0 {
					Fprintln(out, 0)
					continue o
				}
				i--
				targetK = a[i]
			} else if len(s) > 0 && s[len(s)-1].k == k {
				s[len(s)-1].c++
			} else {
				s = append(s, pair{k, 1})
			}
		}
		ans := pow(targetK)
		for _, p := range s {
			ans -= pow(p.k) * int64(p.c)
		}
		Fprintln(out, (ans%mod+mod)%mod)
	}
}

//func main() { CF1361B(os.Stdin, os.Stdout) }
