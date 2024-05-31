package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pre := make([]int, n)
		g := []int{}
		for i := range a {
			Fscan(in, &a[i])
			v := a[i]
			p := sort.SearchInts(g, v)
			if p < len(g) {
				g[p] = v
			} else {
				g = append(g, v)
			}
			pre[i] = p + 1
		}

		suf := make([]int, n)
		g = g[:0]
		for i := n - 1; i >= 0; i-- {
			v := -a[i]
			p := sort.SearchInts(g, v)
			if p < len(g) {
				g[p] = v
			} else {
				g = append(g, v)
			}
			suf[i] = p + 1
		}

		lis := len(g)
		ans := []any{}
		for i, p := range pre {
			if p+suf[i]-1 == lis {
				ans = append(ans, i+1)
			}
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
