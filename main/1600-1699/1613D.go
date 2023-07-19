package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1613D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		f := make([]int64, n+1)
		g2 := make([]int64, n+3)
		l2 := make([]int64, n+1)
		f1 := int64(0)
		for ; n > 0; n-- {
			Fscan(in, &v)
			if v == 0 {
				f[0] = (f[0]*2 + 1) % mod
			} else {
				if v == 1 {
					f1 = (f1*2 + 1) % mod
				} else {
					g2[v] = (g2[v]*2 + l2[v-2] + f[v-2]) % mod
				}
				f[v] = (f[v]*2 + f[v-1]) % mod
			}
			l2[v] = (l2[v]*2 + g2[v+2]) % mod
		}
		ans := f1
		for _, v := range f  { ans += v }
		for _, v := range g2 { ans += v }
		for _, v := range l2 { ans += v }
		Fprintln(out, ans%mod)
	}
}

//func main() { CF1613D(os.Stdin, os.Stdout) }
