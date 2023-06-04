package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1839B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		for ; n > 0; n-- {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
		}
		ans := int64(0)
		for i, a := range g {
			sort.Ints(a)
			if len(a) > i {
				a = a[len(a)-i:]
			}
			for _, v := range a {
				ans += int64(v)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1839B(os.Stdin, os.Stdout) }
