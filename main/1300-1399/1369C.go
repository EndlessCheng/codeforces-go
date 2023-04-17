package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1369C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		w := make([]int, k)
		for i := range w {
			Fscan(in, &w[i])
		}
		sort.Ints(w)

		ans := int64(0)
		for i, c := range w {
			x := int64(a[n-1-i])
			ans += x
			if c == 1 {
				ans += x
			}
		}
		for i, j := k-1, 0; i >= 0 && w[i] > 1; i-- {
			ans += int64(a[j])
			j += w[i] - 1
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1369C(os.Stdin, os.Stdout) }
