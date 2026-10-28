package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1637G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		if n == 2 {
			Fprintln(out, -1)
			continue
		}

		m := 1 << bits.Len(uint(n-1))
		c := make([]int, m*2+1)
		var a, b []int

		op := func(x, y int) {
			a = append(a, x)
			b = append(b, y)
			c[x]--
			c[y]--
			c[x+y]++
			d := x - y
			if d < 0 {
				d = -d
			}
			c[d]++
		}

		var dfs func(int, int)
		dfs = func(x, y int) {
			if y <= 2 {
				return
			}
			z := 1 << (bits.Len(uint(y)) - 1)
			for i := y; i > z; i-- {
				op(x*i, x*(z*2-i))
			}
			dfs(x, z*2-y-1)
			dfs(x*2, y-z)
		}

		for i := 1; i <= n; i++ {
			c[i] = 1
		}
		dfs(1, n)

		for i := 1; i < m; i++ {
			if c[i] > 1 {
				op(i, i)
				break
			}
		}
		for i := 1; i < m; i++ {
			for c[i] != 0 {
				op(0, i)
				op(i, i)
			}
		}
		op(0, m)
		c[m] = 0

		Fprintln(out, len(a))
		for i := range a {
			Fprintln(out, a[i], b[i])
		}
	}
}

//func main() { cf1637G(bufio.NewReader(os.Stdin), os.Stdout) }
