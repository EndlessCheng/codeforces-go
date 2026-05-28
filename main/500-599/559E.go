package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf559E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type pair struct{ p, l int }
	a := make([]pair, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].p, &a[i].l)
	}
	slices.SortFunc(a[1:], func(x, y pair) int { return x.p - y.p })
	a[0].p = -1e9

	f := make([][][2]int, n+1)
	for i := range f {
		f[i] = make([][2]int, n+1)
	}
	ans := 0
	for i := range n + 1 {
		for j := range n + 1 {
			for k := range 2 {
				ans = max(ans, f[i][j][k])
				mx := a[j].p + k*a[j].l
				mx2 := int(-1e9)
				x, y := 0, 0
				for l := i + 1; l <= n; l++ {
					for p := range 2 {
						mx3 := a[l].p + p*a[l].l
						if mx3 > mx2 {
							mx2 = mx3
							x, y = l, p
						}
						f[l][x][y] = max(f[l][x][y], f[i][j][k]+min(a[l].l, mx3-mx)+mx2-mx3)
					}
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf559E(bufio.NewReader(os.Stdin), os.Stdout) }
