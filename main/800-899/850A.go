package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF850A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][5]int, n)
	for i := range a {
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	ans := []int{}
o:
	for i, p := range a {
		for j, q := range a {
			if j == i {
				continue
			}
			for k, r := range a[:j] {
				if k == i {
					continue
				}
				s := 0
				for k, v := range p {
					s += (q[k] - v) * (r[k] - v)
				}
				if s > 0 {
					continue o
				}
			}
		}
		ans = append(ans, i)
	}
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF850A(os.Stdin, os.Stdout) }
