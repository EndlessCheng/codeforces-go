package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1399C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, w int
	for Fscan(in, &T); T > 0; T-- {
		cnt := [51]int{}
		for Fscan(in, &n); n > 0; n-- {
			Fscan(in, &w)
			cnt[w]++
		}
		ans := 0
		for s := 2; s < 101; s++ {
			c := 0
			for i := max(1, s-50); 2*i <= s; i++ {
				if 2*i < s {
					c += min(cnt[i], cnt[s-i])
				} else {
					c += cnt[i] / 2
				}
			}
			ans = max(ans, c)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1399C(os.Stdin, os.Stdout) }
