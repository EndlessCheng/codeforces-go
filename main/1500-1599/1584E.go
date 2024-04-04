package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1584E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ s, cnt int }
		q := make([]pair, n+1)
		q[n/2].cnt = 1 // pair{0, 1}
		l, r := n/2, n/2
		ans, s := 0, 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if i%2 == 0 {
				s += v
				for l <= r && q[r].s > s {
					r--
				}
				if l <= r && q[r].s == s {
					ans += q[r].cnt
					q[r].cnt++
				} else {
					r++
					q[r] = pair{s, 1}
				}
			} else {
				s -= v
				for l <= r && q[l].s < s {
					l++
				}
				if l <= r && q[l].s == s {
					ans += q[l].cnt
					q[l].cnt++
				} else {
					l--
					q[l] = pair{s, 1}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1584E(os.Stdin, os.Stdout) }
