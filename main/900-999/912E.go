package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf912E(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n)
	p := make([]int, n)
	for i := range p {
		Fscan(in, &p[i])
	}
	Fscan(in, &k)

	gen := func(st int) (a []int) {
		var f func(int, int)
		f = func(i, m int) {
			if i >= len(p) {
				a = append(a, m)
				return
			}
			f(i+2, m)
			if m <= 1e18/p[i] {
				f(i, m*p[i])
			}
		}
		f(st, 1)
		slices.Sort(a)
		return
	}
	a := gen(0)
	b := gen(1)

	ans := sort.Search(1e18, func(mx int) bool {
		k := k
		j := len(b) - 1
		for _, v := range a {
			for j >= 0 && b[j] > mx/v {
				j--
			}
			k -= j + 1
			if k <= 0 {
				return true
			}
		}
		return false
	})
	Fprint(out, ans)
}

//func main() { cf912E(os.Stdin, os.Stdout) }
