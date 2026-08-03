package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf55E(in io.Reader, out io.Writer) {
	type vec struct{ x, y int }
	cross := func(a, b, c vec) int { return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x) }

	var n, q int
	var p vec
	Fscan(in, &n)
	a := make([]vec, n+1)
	for i := range n {
		Fscan(in, &a[i].x, &a[i].y)
	}
	a[n] = a[0]

	Fscan(in, &q)
	for range q {
		Fscan(in, &p.x, &p.y)
		ans := n * (n - 1) * (n - 2) / 6
		j := 1
		for i := range n {
			if cross(p, a[i], a[i+1]) > 0 {
				ans = 0
				break
			}
			for cross(p, a[i], a[j]) < 0 {
				j = (j + 1) % n
			}
			k := (n + i - j - 1) % n
			ans -= k * (k + 1) / 2
		}
		Fprintln(out, ans)
	}
}

//func main() { cf55E(bufio.NewReader(os.Stdin), os.Stdout) }
