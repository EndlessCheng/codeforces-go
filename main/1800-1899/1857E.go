package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1857E(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ x, i int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x)
			a[i].i = i
		}
		slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })

		s := n
		for _, p := range a {
			s += p.x - a[0].x
		}

		ans := make([]any, n)
		ans[a[0].i] = s
		for i := 1; i < n; i++ {
			s -= (n - i*2) * (a[i].x - a[i-1].x)
			ans[a[i].i] = s
		}
		Fprintln(out, ans...)
	}
}

//func main() { cf1857E(bufio.NewReader(os.Stdin), os.Stdout) }
