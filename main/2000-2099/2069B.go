package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2069B(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n)
		c := make([]int8, n*m+1)
		c1, c2 := 0, 0
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
				v := a[i][j]
				if c[v] == 2 {
					continue
				}
				if j > 0 && v == a[i][j-1] || i > 0 && v == a[i-1][j] {
					c[v] = 2
					c2++
					c1--
				} else if c[v] == 0 {
					c[v] = 1
					c1++
				}
			}
		}
		if c2 == 0 {
			Fprintln(out, c1-1)
		} else {
			Fprintln(out, c1+(c2-1)*2)
		}
	}
}

//func main() { cf2069B(bufio.NewReader(os.Stdin), os.Stdout) }
