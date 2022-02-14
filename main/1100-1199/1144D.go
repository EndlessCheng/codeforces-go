package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1144D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, mx, mxC int
	Fscan(in, &n)
	a := make([]int, n)
	c := [2e5 + 1]int{}
	for i := range a {
		Fscan(in, &v)
		a[i] = v
		if c[v]++; c[v] > mxC {
			mx, mxC = v, c[v]
		}
	}
	Fprintln(out, n-mxC)
	for i, v := range a {
		if v != mx {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			op := 1
			if a[j] > mx {
				op = 2
			}
			Fprintln(out, op, j+1, j+2)
		}
		for j := i + 1; j < n; j++ {
			if a[j] != mx {
				op := 1
				if a[j] > mx {
					op = 2
				}
				Fprintln(out, op, j+1, j)
			}
		}
		break
	}
}

//func main() { CF1144D(os.Stdin, os.Stdout) }
