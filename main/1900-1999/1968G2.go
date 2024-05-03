package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1968G2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l, r int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r, &s)
		z := make([]int, n)
		bl, br := 0, 0
		for i := 1; i < n; i++ {
			if i <= br {
				z[i] = min(z[i-bl], br-i+1)
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				bl, br = i, i+z[i]
				z[i]++
			}
		}

		memo := make([]int, n+1)
		getK := func(lcp int) int {
			if memo[lcp] > 0 {
				return memo[lcp]
			}
			cnt := 1
			for i := lcp; i <= n-lcp; {
				if z[i] >= lcp {
					cnt++
					i += lcp
				} else {
					i++
				}
			}
			memo[lcp] = cnt
			return cnt
		}
		for i := l; i <= r; i++ {
			Fprint(out, sort.Search(n/i, func(lcp int) bool { return getK(lcp+1) < i }), " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1968G2(os.Stdin, os.Stdout) }
