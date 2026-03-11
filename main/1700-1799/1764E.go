package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1764E(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		type pair struct{ a, b int }
		e := make([]pair, n)
		for i := range e {
			Fscan(in, &e[i].a, &e[i].b)
		}
		slices.SortFunc(e[1:], func(x, y pair) int { return x.a + x.b - y.a - y.b })

		f := int(-1e9)
		mx := 0
		for i := 1; i < n; i++ {
			a, b := e[i].a, e[i].b
			f = max(mx, max(a, min(f, a)+b))
			mx = max(mx, a+b)
		}

		a, b := e[0].a, e[0].b
		if max(a, min(f, a)+b) >= k {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1764E(bufio.NewReader(os.Stdin), os.Stdout) }
