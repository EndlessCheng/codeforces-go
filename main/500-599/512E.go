package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf512E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([][]bool, n+1)
	b := make([][]bool, n+1)
	for i := range a {
		a[i] = make([]bool, n+1)
		b[i] = make([]bool, n+1)
	}
	for range n - 3 {
		var x, y int
		Fscan(in, &x, &y)
		a[x][y] = true
		a[y][x] = true
	}
	for range n - 3 {
		var x, y int
		Fscan(in, &x, &y)
		b[x][y] = true
		b[y][x] = true
	}

	qx := make([]int, 0, n-3)
	qy := make([]int, 0, n-3)
	sx := make([]int, 0, n-3)
	for i := 2; i <= n; i++ {
		for j := n; j >= i+2; j-- {
			if a[i][j] {
				qx = append(qx, i)
				qy = append(qy, j)
				a[i][j] = false
				a[j][i] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		for j := n; j >= i+2; j-- {
			if b[i][j] {
				k := j - 1
				for k > i {
					if b[i][k] || k == i+1 {
						break
					}
					k--
				}
				sx = append(sx, k)
				b[i][j] = false
				b[j][i] = false
				b[1][k] = true
				b[k][1] = true
			}
		}
	}

	Fprintln(out, len(qx)+len(sx))
	for i := range qx {
		Fprintln(out, qx[i], qy[i])
	}
	for i := len(sx) - 1; i >= 0; i-- {
		Fprintln(out, 1, sx[i])
	}
}

//func main() { cf512E(bufio.NewReader(os.Stdin), os.Stdout) }
