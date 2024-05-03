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
		getK := func(lcp int) int {
			cnt := 1
			for i := lcp; i <= n-lcp; {
				if z[i] >= lcp {
					cnt++
					i += lcp
				} else {
					i++
				}
			}
			return cnt
		}

		f := make([]int, n+1)
		i := r
		for lcp := 1; lcp*lcp <= n; lcp++ {
			// 给定 LCP，最多分多少段？
			for k := getK(lcp); i > k; i-- {
				f[i] = lcp - 1
			}
		}
		for ; i >= l; i-- {
			// 给定 i，分成 i 段，LCP 是多少？
			f[i] = sort.Search(n/i, func(lcp int) bool { return getK(lcp+1) < i })
		}

		for _, v := range f[l : r+1] {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1968G2(os.Stdin, os.Stdout) }
