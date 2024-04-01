package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1538C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &l, &r)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		ans := 0
		for i, v := range a {
			j := sort.SearchInts(a[:i], l-v)
			k := sort.SearchInts(a[:i], r-v+1)
			ans += k - j
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1538C(os.Stdin, os.Stdout) }
