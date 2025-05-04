package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p1496(in io.Reader, out io.Writer) {
	var n, l, r, s, ans int
	d := map[int]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &l, &r)
		d[l]++
		d[r]--
	}

	a := make([]int, 0, len(d))
	for k := range d {
		a = append(a, k)
	}
	slices.Sort(a)
	for i, v := range a {
		if s > 0 {
			ans += v - a[i-1]
		}
		s += d[v]
	}
	Fprint(out, ans)
}

//func main() { p1496(bufio.NewReader(os.Stdin), os.Stdout) }
