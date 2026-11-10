package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3069(in io.Reader, out io.Writer) {
	var n, k, l, maxCnt, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	cnt := map[int]int{}
	cc := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		cnt[a[i]]++
		c := cnt[a[i]]
		cc[c-1]--
		cc[c]++
		maxCnt = max(maxCnt, c)
		for len(cnt) > k+1 {
			v := a[l]
			cnt[v]--
			c := cnt[v]
			if c == 0 {
				delete(cnt, v)
			}
			cc[c+1]--
			cc[c]++
			for cc[maxCnt] == 0 {
				maxCnt--
			}
			l++
		}
		ans = max(ans, maxCnt)
	}
	Fprint(out, ans)
}

//func main() { p3069(bufio.NewReader(os.Stdin), os.Stdout) }
