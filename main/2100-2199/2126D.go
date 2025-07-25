package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2126D(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		type pair struct{ l, real int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].l, &n, &a[i].real)
		}
		slices.SortFunc(a, func(a, b pair) int { return a.real - b.real })
		for _, t := range a {
			if t.l <= k && k < t.real {
				k = t.real
			}
		}
		Fprintln(out, k)
	}
}

//func main() { cf2126D(bufio.NewReader(os.Stdin), os.Stdout) }
