package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1181E1(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type tuple struct{ u, d, l, r int }
	a := make([]tuple, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].l, &a[i].d, &a[i].r, &a[i].u)
	}

	var solve func(int, int) bool
	solve = func(left, right int) bool {
		if left == right {
			return true
		}

		slices.SortFunc(a[left:right+1], func(x, y tuple) int { return cmp.Or(x.l-y.l, x.r-y.r) })

		p := -1
		mx := 0
		for i := left + 1; i <= right; i++ {
			mx = max(mx, a[i-1].r)
			if mx <= a[i].l {
				p = i
				break
			}
		}
		if p != -1 {
			return solve(left, p-1) && solve(p, right)
		}

		slices.SortFunc(a[left:right+1], func(x, y tuple) int { return cmp.Or(x.d-y.d, x.u-y.u) })

		mx = 0
		for i := left + 1; i <= right; i++ {
			mx = max(mx, a[i-1].u)
			if mx <= a[i].d {
				p = i
				break
			}
		}
		if p != -1 {
			return solve(left, p-1) && solve(p, right)
		}
		return false
	}

	if solve(1, n) {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { cf1181E1(bufio.NewReader(os.Stdin), os.Stdout) }
