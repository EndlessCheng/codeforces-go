package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1623C(_r io.Reader, _w io.Writer) {
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
		ans := sort.Search(1e9, func(low int) bool {
			low++
			inc, inc2 := 0, 0
			for i := n - 1; i > 1; i-- {
				d := a[i] + inc - low
				if d < 0 {
					return true
				}
				d = min(d, a[i])
				inc = inc2 + d/3
				inc2 = d / 3 * 2
			}
			return a[1]+inc < low || a[0]+inc2 < low
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1623C(os.Stdin, os.Stdout) }
