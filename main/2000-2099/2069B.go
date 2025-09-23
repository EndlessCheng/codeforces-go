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
		ans := -1
		has2 := 0
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
					has2 = 1
					ans++
				} else if c[v] == 0 {
					c[v] = 1
					ans++
				}
			}
		}
		Fprintln(out, ans-has2)
	}
}

//func main() { cf2069B(bufio.NewReader(os.Stdin), os.Stdout) }
