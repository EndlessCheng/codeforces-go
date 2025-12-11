package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1333E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	if n < 3 {
		Fprint(out, -1)
		return
	}

	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}

	b := n*n - 9
	a[0][0] = b + 1
	a[0][1] = b + 2
	a[0][2] = b + 4
	a[1][0] = b + 5
	a[1][1] = b + 3
	a[1][2] = b + 8
	a[2][0] = b + 9
	a[2][1] = b + 6
	a[2][2] = b + 7

	v := 1
	for i := 3; i < n; i++ {
		if i%2 > 0 {
			for j := range i + 1 {
				a[i][j] = v
				v++
			}
			for j := i - 1; j >= 0; j-- {
				a[j][i] = v
				v++
			}
		} else {
			for j := range i + 1 {
				a[j][i] = v
				v++
			}
			for j := i - 1; j >= 0; j-- {
				a[i][j] = v
				v++
			}
		}
	}

	for _, r := range a {
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1333E(bufio.NewReader(os.Stdin), os.Stdout) }
