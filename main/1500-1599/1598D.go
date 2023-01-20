package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1598D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y int }, n)
		cx := make([]int64, n+1)
		cy := make([]int64, n+1)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
			cx[a[i].x]++
			cy[a[i].y]++
		}
		ans := n * (n - 1) * (n - 2) / 6
		for _, p := range a {
			ans -= (cx[p.x] - 1) * (cy[p.y] - 1)
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1598D(os.Stdin, os.Stdout) }
