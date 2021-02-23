package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1492E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	var f func(int) bool
	f = func(canModify int) bool {
		for _, r := range a[1:] {
			d := []int{}
			for j, v := range r {
				if v != a[0][j] {
					d = append(d, j)
				}
			}
			if len(d) <= 2 {
				continue
			}
			if len(d) <= 2+canModify {
				for _, j := range d {
					tmp := a[0][j]
					a[0][j] = r[j]
					if f(canModify - 1) {
						return true
					}
					a[0][j] = tmp
				}
			}
			return false
		}
		return true
	}
	if f(2) {
		Fprintln(out, "Yes")
		for _, v := range a[0] {
			Fprint(out, v, " ")
		}
	} else {
		Fprint(out, "No")
	}
}

//func main() { CF1492E(os.Stdin, os.Stdout) }
