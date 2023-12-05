package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1027C(_r io.Reader, _w io.Writer) {
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
		x, y := 1, 0
		pre := 0
		for i := 1; i < n; i++ {
			if a[i] != a[i-1] {
				continue
			}
			if a[i]*y < pre*x {
				x, y = a[i], pre
			}
			pre = a[i]
			i++
		}
		Fprintln(out, x, x, y, y)
	}
}

//func main() { cf1027C(os.Stdin, os.Stdout) }
