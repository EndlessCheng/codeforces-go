package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1166E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var m, n, v int
	Fscan(in, &m, &n)
	a := make([][]bool, m)
	for i := range a {
		a[i] = make([]bool, n+1)
		for Fscan(in, &m); m > 0; m-- {
			Fscan(in, &v)
			a[i][v] = true
		}
	}
	for i, vs := range a {
	o:
		for _, ws := range a[i+1:] {
			for j, v := range vs {
				if v && ws[j] {
					continue o
				}
			}
			Fprint(out, "impossible")
			return
		}
	}
	Fprint(out, "possible")
}

//func main() { CF1166E(os.Stdin, os.Stdout) }
