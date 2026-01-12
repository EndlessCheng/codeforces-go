package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1458C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		mat := make([][]int, n)
		for i := range mat {
			mat[i] = make([]int, n)
			for j := range mat[i] {
				Fscan(in, &mat[i][j])
			}
		}
		Fscan(in, &s)

		x, y, z := 0, 1, 2
		a, b, c := 0, 0, 0
		for _, ch := range s {
			switch ch {
			case 'R':
				b = (b + 1) % n
			case 'L':
				b = (b - 1 + n) % n
			case 'D':
				a = (a + 1) % n
			case 'U':
				a = (a - 1 + n) % n
			case 'I':
				b, c = c, b
				y, z = z, y
			case 'C':
				a, c = c, a
				x, z = z, x
			}
		}

		ans := make([][]int, n)
		for i := range ans {
			ans[i] = make([]int, n)
		}
		for i := range n {
			for j := range n {
				p := [3]int{i, j, mat[i][j] - 1}
				ans[(p[x]+a)%n][(p[y]+b)%n] = (p[z]+c)%n + 1
			}
		}

		for _, r := range ans {
			for _, v := range r {
				Fprint(out, v, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf1458C(bufio.NewReader(os.Stdin), os.Stdout) }
