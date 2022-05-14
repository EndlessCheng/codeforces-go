package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF463C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][]int, n)
	b := make([]int64, n*2)
	c := make([]int64, n*2)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
			b[i-j+n] += int64(a[i][j])
			c[i+j] += int64(a[i][j])
		}
	}
	mx := [2]int64{-1, -1}
	p := [2][2]int{}
	for i, r := range a {
		for j, v := range r {
			if s, k := b[i-j+n]+c[i+j]-int64(v), (i+j)&1; s > mx[k] {
				mx[k], p[k] = s, [2]int{i, j}
			}
		}
	}
	Fprintln(out, mx[0]+mx[1])
	Fprintln(out, p[0][0]+1, p[0][1]+1, p[1][0]+1, p[1][1]+1)
}

//func main() { CF463C(os.Stdin, os.Stdout) }
