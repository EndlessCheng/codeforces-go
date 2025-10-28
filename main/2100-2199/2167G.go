package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2167G(in io.Reader, out io.Writer) {
	var T, n, c int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		f := make([]int, n)
		s := 0
		for i, v := range a {
			Fscan(in, &c)
			s += c
			for j, w := range a[:i] {
				if w <= v {
					f[i] = max(f[i], f[j])
				}
			}
			f[i] += c
		}
		Fprintln(out, s-slices.Max(f))
	}
}

//func main() { cf2167G(bufio.NewReader(os.Stdin), os.Stdout) }
