package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1450C1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][]byte, n)
		c := [3]int{}
		for i := range a {
			Fscan(in, &a[i])
			for j, b := range a[i] {
				if b == 'X' {
					c[(i+j)%3]++
				}
			}
		}
		miI := 0
		if c[1] < c[0] {
			miI = 1
		}
		if c[2] < c[miI] {
			miI = 2
		}
		for i, r := range a {
			for j, b := range r {
				if b == 'X' && (i+j)%3 == miI {
					r[j] = 'O'
				}
			}
			Fprintln(out, string(r))
		}
	}
}

//func main() { CF1450C1(os.Stdin, os.Stdout) }
