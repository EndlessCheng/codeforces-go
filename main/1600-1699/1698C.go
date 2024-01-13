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
		for i := range a {
			Fscan(in, &a[i])
		}
		sort.Ints(a)
		if a[2] < 0 || a[1] < 0 && a[2] == 0 || a[n-3] > 0 || a[n-2] > 0 && a[n-3] == 0 {
			Fprintln(out, "NO")
			continue
		}
		has := map[int]bool{}
		for _, v := range a {
			has[v] = true
		}
		if n == 3 && !has[a[0]+a[1]+a[2]] ||
			n == 4 && (!has[a[0]+a[1]+a[2]] || !has[a[0]+a[1]+a[3]] || !has[a[0]+a[2]+a[3]] || !has[a[1]+a[2]+a[3]]) ||
			n > 4 && !has[a[0]+a[n-1]] {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { cf1698C(os.Stdin, os.Stdout) }
