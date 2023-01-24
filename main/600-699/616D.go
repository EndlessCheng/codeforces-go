package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF616D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, c, rr int
	ll := 1
	Fscan(in, &n, &k)
	cnt := [1e6 + 1]int{}
	a := make([]int, n+1)
	for l, r := 1, 1; r <= n; r++ {
		Fscan(in, &a[r])
		if cnt[a[r]] == 0 {
			c++
		}
		cnt[a[r]]++
		for c > k {
			cnt[a[l]]--
			if cnt[a[l]] == 0 {
				c--
			}
			l++
		}
		if r-l > rr-ll {
			ll, rr = l, r
		}
	}
	Fprint(out, ll, rr)
}

//func main() { CF616D(os.Stdin, os.Stdout) }
