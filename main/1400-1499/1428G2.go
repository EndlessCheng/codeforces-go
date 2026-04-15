package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1428G2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var k int
	Fscan(in, &k)
	f := [6]int{}
	for i := range f {
		Fscan(in, &f[i])
	}

	g := [1e6]int{}
	for i := range len(g) {
		j := 0
		for x := i; x > 0; x /= 10 {
			d := x % 10
			if d%3 == 0 {
				g[i] += d / 3 * f[j]
			}
			j++
		}
	}

	p10 := 3
	for i := range f {
		q := 3 * (k - 1)
		p2 := 1
		for p2 < q {
			for x := len(g) - 1; x >= p2*p10; x-- {
				if g[x] < g[x-p2*p10]+f[i]*p2 {
					g[x] = g[x-p2*p10] + f[i]*p2
				}
			}
			q -= p2
			p2 *= 2
		}
		for x := len(g) - 1; x >= q*p10; x-- {
			if g[x] < g[x-q*p10]+f[i]*q {
				g[x] = g[x-q*p10] + f[i]*q
			}
		}
		p10 *= 10
	}

	var q int
	Fscan(in, &q)
	for range q {
		var x int
		Fscan(in, &x)
		Fprintln(out, g[x])
	}
}

//func main() { cf1428G2(bufio.NewReader(os.Stdin), os.Stdout) }
