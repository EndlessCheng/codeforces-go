package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf13D(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	x := make([]int, n)
	y := make([]int, n)
	x2 := make([]int, m)
	y2 := make([]int, m)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range n {
		Fscan(in, &x[i], &y[i])
	}
	for i := range m {
		Fscan(in, &x2[i], &y2[i])
	}

	for i := range n {
		for j := range n {
			if x[i] < x[j] {
				a := y[j] - y[i]
				b := x[i] - x[j]
				c := -x[i]*a - y[i]*b
				for k := range m {
					if x2[k] > x[i] && x2[k] <= x[j] && a*x2[k]+b*y2[k]+c > 0 {
						f[i][j]++
					}
				}
				f[j][i] = -f[i][j]
			}
		}
	}

	ans := 0
	for i := range n {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if f[i][j]+f[j][k]+f[k][i] == 0 {
					ans++
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf13D(bufio.NewReader(os.Stdin), os.Stdout) }
