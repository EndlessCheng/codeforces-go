package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1932F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		cnt := make([]int, n+2)
		rs := make([][]int, n+1)
		for ; m > 0; m-- {
			Fscan(in, &l, &r)
			cnt[l]++
			cnt[r+1]--
			rs[l] = append(rs[l], r)
		}

		nxt := make([]int, n+1)
		mx := 0
		for i := 1; i <= n; i++ {
			cnt[i] += cnt[i-1]
			for _, r := range rs[i] {
				mx = max(mx, r+1)
			}
			mx = max(mx, i+1)
			nxt[i] = mx
		}

		f := make([]int, n+2)
		for i := n; i > 0; i-- {
			f[i] = max(f[i+1], f[nxt[i]]+cnt[i])
		}
		Fprintln(out, f[1])
	}
}

//func main() { cf1932F(os.Stdin, os.Stdout) }
