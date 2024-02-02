package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1914E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ x, y int }, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
		}
		sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.x+a.y > b.x+b.y })
		d := 0
		for i, p := range a {
			if i%2 == 0 {
				d += p.x
			} else {
				d -= p.y
			}
		}
		Fprintln(out, d-n%2)
	}
}

//func main() { cf1914E2(os.Stdin, os.Stdout) }
