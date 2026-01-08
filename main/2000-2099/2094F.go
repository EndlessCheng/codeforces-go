package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2094F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
		}
		v, c := 0, 0
		f := func(i, j int) {
			if c == 0 {
				v++
				c = n * m / k
			}
			a[i][j] = v
			c--
		}
		for i := range n {
			for j := i % 2; j < m; j += 2 {
				f(i, j)
			}
		}
		for i := range n {
			for j := i%2 ^ 1; j < m; j += 2 {
				f(i, j)
			}
		}
		for _, r := range a {
			for _, v := range r {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf2094F(bufio.NewReader(os.Stdin), os.Stdout) }
