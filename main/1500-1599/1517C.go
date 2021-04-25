package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1517C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, i+1)
		Fscan(in, &a[i][i])
	}
	for i := 1; i < n; i++ {
		f := false
		for j := 0; j+i < n; j++ {
			if a[j+i-1][j] == i {
				f = true
			}
			if f {
				a[j+i][j] = a[j+i][j+1]
			} else {
				a[j+i][j] = a[j+i-1][j]
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

//func main() { CF1517C(os.Stdin, os.Stdout) }
