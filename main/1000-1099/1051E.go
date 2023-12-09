package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func calcZ51(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}

func cf1051E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	var S, L, R string
	Fscan(in, &S, &L, &R)

	n, nl, nr := len(S), len(L), len(R)
	if nl > n {
		Fprint(out, 0)
		return
	}
	zl := calcZ51(L + S)
	zr := calcZ51(R + S)

	sum := make([]int, n+2)
	sum[1] = 1
	pre := 0
	for i := 1; i <= n; i++ {
		sum[i+1] = sum[i]
		if i < nl {
			pre = 0
			continue
		}
		
		// [l,r]
		l, r := i-nr, i-nl
		if lcp := zl[nl+r]; lcp < nl && S[r+lcp] < L[lcp] {
			r--
			if r < 0 {
				pre = 0
				continue
			}
		}
		if l < 0 {
			l = 0
		} else {
			if lcp := zr[nr+l]; lcp < nr && S[l+lcp] > R[lcp] {
				l++
			}
		}
		if l > r {
			pre = 0
			continue
		}
		
		f := 0
		if L == "0" && S[i-1] == '0' {
			f = pre
			r--
		}
		f = (f + sum[r+1] - sum[l]) % mod
		pre = f
		if i < n && S[i] == '0' {
			f = 0
		}
		sum[i+1] += f
	}
	Fprint(out, pre)
}

//func main() { cf1051E(os.Stdin, os.Stdout) }
