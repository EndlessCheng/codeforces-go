package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1776D(in io.Reader, out io.Writer) {
	var l int
	w := [5]int{}
	t := [3]int{}
	Fscan(in, &w[2], &w[3], &w[4], &l)
	x := make([]int, l)
	p := make([]int, l)
	q := make([]int, l)
	n := 0

	for i := 1; i <= l; i++ {
		for k := 4; k > 1; k-- {
			if w[k] != 0 {
				j := -1
				for o := 0; o < 3; o++ {
					if t[o]+k <= i && (j == -1 || t[o] > t[j]) {
						j = o
					}
				}
				if j != -1 {
					x[n] = j + 1
					p[n] = i - k
					q[n] = i
					n++
					w[k]--
					t[j] = i
					break
				}
			}
		}
	}

	Fprintln(out, n)
	for i := 0; i < n; i++ {
		Fprintln(out, x[i], p[i], q[i])
	}
}

//func main() { cf1776D(bufio.NewReader(os.Stdin), os.Stdout) }
