package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2248B(in io.Reader, out io.Writer) {
	var T, n, m int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		if n < m*2 {
			Fprintln(out, "NO")
			continue
		}

		slices.Sort(a)
		slices.Sort(b)

		for i, v := range b {
			if v < a[i] || v > a[n-m+i] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf2248B(bufio.NewReader(os.Stdin), os.Stdout) }
