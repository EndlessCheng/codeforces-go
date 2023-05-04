package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1774B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, k int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make(sort.IntSlice, m)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Sort(sort.Reverse(a))
		for i, v := range a {
			limit := n / k
			if i < n%k {
				limit++
			}
			if v > limit {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1774B(os.Stdin, os.Stdout) }
