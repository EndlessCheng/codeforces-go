package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf535E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type vec struct{ x, y int }
	det := func(c, a, b vec) bool { return (b.x*a.x-a.x*c.x)*(b.y*c.y-a.y*c.y)-(b.x*c.x-a.x*c.x)*(a.y*b.y-a.y*c.y) < 0 }

	var n int
	Fscan(in, &n)
	a := make([]vec, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	ori := slices.Clone(a)

	slices.SortFunc(a, func(a, b vec) int { return cmp.Or(b.x-a.x, b.y-a.y) })
	c := make([]vec, n)
	c[0] = a[0]
	m := 1
	for i := 1; i < n; i++ {
		if a[i].y <= c[m-1].y {
			continue
		}
		for m >= 2 && det(c[m-2], c[m-1], a[i]) {
			m--
		}
		c[m] = a[i]
		m++
	}

	t := map[vec]bool{}
	for i := range m {
		t[c[i]] = true
	}
	for i, v := range ori {
		if t[v] {
			Fprint(out, i+1, " ")
		}
	}
}

//func main() { cf535E(bufio.NewReader(os.Stdin), os.Stdout) }
