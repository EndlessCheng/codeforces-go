package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1624E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		type tuple struct{ l, r, i int }
		mp := map[string]tuple{}
		Fscan(in, &n, &m)
		for i := 1; i <= n; i++ {
			Fscan(in, &s)
			for r := 2; r <= m; r++ {
				mp[s[r-2:r]] = tuple{r - 1, r, i}
				if r > 2 {
					mp[s[r-3:r]] = tuple{r - 2, r, i}
				}
			}
		}

		Fscan(in, &s)
		f := make([]bool, m+1)
		f[0] = true
		for i := 2; i <= m; i++ {
			f[i] = f[i-2] && mp[s[i-2:i]].l > 0 || i > 2 && f[i-3] && mp[s[i-3:i]].l > 0
		}
		if !f[m] {
			Fprintln(out, -1)
			continue
		}

		ans := []tuple{}
		for i := m; i > 0; {
			if f[i-2] && mp[s[i-2:i]].l > 0 {
				ans = append(ans, mp[s[i-2:i]])
				i -= 2
			} else {
				ans = append(ans, mp[s[i-3:i]])
				i -= 3
			}
		}
		Fprintln(out, len(ans))
		for i := len(ans) - 1; i >= 0; i-- {
			t := ans[i]
			Fprintln(out, t.l, t.r, t.i)
		}
	}
}

//func main() { CF1624E(os.Stdin, os.Stdout) }
