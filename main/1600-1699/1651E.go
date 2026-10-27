package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1651E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([][]int, n+1)
	for range 2 * n {
		var x, y int
		Fscan(in, &x, &y)
		a[x] = append(a[x], y-n)
	}

	fa := [2][]int{make([]int, n+1), make([]int, n+1)}

	var find func(int, int) int
	find = func(t, x int) int {
		if fa[t][x] != x {
			fa[t][x] = find(t, fa[t][x])
		}
		return fa[t][x]
	}

	ans := 0
	for i := 1; i <= n; i++ {
		for j := range fa {
			for k := 1; k <= n; k++ {
				fa[j][k] = k
			}
		}

		s := 0
		for j := i; j <= n; j++ {
			la := find(0, a[j][0])
			ra := find(1, a[j][0])
			lb := find(0, a[j][1])
			rb := find(1, a[j][1])
			s += la*(n+1-ra) + lb*(n+1-rb)

			if la > lb {
				la, lb = lb, la
			}
			if ra < rb {
				ra, rb = rb, ra
			}

			fa[0][lb] = la
			fa[1][rb] = ra
			s -= la * (n - ra + 1)
			ans += s
		}
	}
	Fprint(out, ans)
}

//func main() { cf1651E(bufio.NewReader(os.Stdin), os.Stdout) }
