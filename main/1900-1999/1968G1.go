package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1968G1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &l, &s)
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
		ans := sort.Search(n/l, func(lcp int) bool { return getK(lcp+1) < l })
		Fprintln(out, ans)
	}
}

//func main() { cf1968G1(os.Stdin, os.Stdout) }
