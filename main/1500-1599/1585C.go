package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1585C(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)

		ans := 0
		p := sort.SearchInts(a, 0)
		for i := 0; i < p; i += k {
			ans -= a[i]
		}
		for i := n - 1; i >= p; i -= k {
			ans += a[i]
		}
		Fprintln(out, ans*2-max(-a[0], a[n-1]))
	}
}

//func main() { cf1585C(bufio.NewReader(os.Stdin), os.Stdout) }
