package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1878E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, l, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		Fscan(in, &q)
		type pair struct{ k, i int }
		qs := make([][]pair, n)
		for i := range q {
			Fscan(in, &l, &k)
			qs[l-1] = append(qs[l-1], pair{k, i})
		}

		ans := make([]int, q)
		for l := n - 1; l >= 0; l-- {
			v := a[l]
			for r := l + 1; r < n && a[r]&v != a[r]; r++ {
				a[r] &= v
			}
			for _, p := range qs[l] {
				r := sort.Search(n-l, func(i int) bool { return a[l+i] < p.k })
				if r > 0 {
					r += l
				} else {
					r = -1
				}
				ans[p.i] = r
			}
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1878E(bufio.NewReader(os.Stdin), os.Stdout) }
