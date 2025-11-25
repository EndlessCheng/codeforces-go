package main

import (
	"bytes"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2053B(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ l, r int }
		a := make([]pair, n)
		cnt := make([]int, n*2+1)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
			if a[i].l == a[i].r {
				cnt[a[i].l]++
			}
		}

		// 优化前 https://codeforces.com/problemset/submission/2053/350549792
		s := make([]int, n*2+1)
		for i := 1; i <= n*2; i++ {
			s[i] = s[i-1]
			if cnt[i] == 0 {
				s[i]--
			}
		}

		ans := bytes.Repeat([]byte{'0'}, n)
		for i, p := range a {
			if p.l < p.r {
				if s[p.r] < s[p.l-1] {
					ans[i] = '1'
				}
			} else if cnt[p.l] == 1 {
				ans[i] = '1'
			}
		}
		Fprintf(out, "%s\n", ans)
	}
}

//func main() { cf2053B(bufio.NewReader(os.Stdin), os.Stdout) }
