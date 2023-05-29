package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1830B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y int }, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
		}
		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })
		ans := int64(0)
		for s := 1; s*s <= n*2; s++ {
			cnt := make([]int, n+1)
			for _, p := range a {
				v := p.x*s - p.y
				if 1 <= v && v <= n {
					ans += int64(cnt[v])
				}
				if p.x == s {
					cnt[p.y]++
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1830B(os.Stdin, os.Stdout) }
