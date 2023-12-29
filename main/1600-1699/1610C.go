package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1610C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y int }, n)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
		}
		ans := sort.Search(n, func(m int) bool {
			m++
			cnt := 0
			for _, p := range a {
				if cnt <= p.y && m-1-cnt <= p.x {
					cnt++
				}
			}
			return cnt < m
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1610C(os.Stdin, os.Stdout) }
