package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf682E(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, s int
	Fscan(in, &n, &s)
	x := make([]int, n+1)
	y := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &x[i], &y[i])
	}

	area := func(a, b, c int) int {
		return abs((x[b]-x[a])*(y[c]-y[a]) - (y[b]-y[a])*(x[c]-x[a]))
	}

	a, b, c := 1, 2, 3
	for flg := 0; flg == 0; {
		flg = 1
		for i := 1; i <= n; i++ {
			if area(i, b, c) > area(a, b, c) {
				a = i
				flg = 0
			}
			if area(a, i, c) > area(a, b, c) {
				b = i
				flg = 0
			}
			if area(a, b, i) > area(a, b, c) {
				c = i
				flg = 0
			}
		}
	}

	Fprintln(out, x[b]+x[c]-x[a], y[b]+y[c]-y[a])
	Fprintln(out, x[a]+x[c]-x[b], y[a]+y[c]-y[b])
	Fprintln(out, x[b]+x[a]-x[c], y[b]+y[a]-y[c])
}

//func main() { cf682E(bufio.NewReader(os.Stdin), os.Stdout) }
