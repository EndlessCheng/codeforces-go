package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf248E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q int
	Fscan(in, &n)
	a := make([]int, n)
	f := make([][102]float64, n)
	for i := range a {
		Fscan(in, &a[i])
		f[i][a[i]] = 1
	}
	oriA := slices.Clone(a)

	ans := 0.
	for i := range f {
		ans += f[i][0]
	}

	Fscan(in, &q)
	for range q {
		var v, w, k int
		Fscan(in, &v, &w, &k)
		v--
		w--
		ans -= f[v][0]
		for range k {
			for j := range oriA[v] + 1 {
				f[v][j] = f[v][j]*float64(a[v]-j)/float64(a[v]) + f[v][j+1]*float64(j+1)/float64(a[v])
			}
			a[v]--
		}
		ans += f[v][0]
		Fprintf(out, "%.9f\n", ans)
		a[w] += k
	}
}

//func main() { cf248E(bufio.NewReader(os.Stdin), os.Stdout) }
