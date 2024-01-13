package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1698C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		has := map[int]bool{}
		for i := range a {
			Fscan(in, &a[i])
			has[a[i]] = true
		}
		sort.Ints(a)
		if has[a[0]+a[1]+a[2]] && has[a[0]+a[1]+a[n-1]] && has[a[0]+a[n-2]+a[n-1]] && has[a[n-3]+a[n-2]+a[n-1]] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1698C(os.Stdin, os.Stdout) }
