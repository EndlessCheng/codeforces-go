package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1450C2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([][]byte, n)
		c := [2][3]int{}
		for i := range a {
			Fscan(in, &a[i])
			for j, b := range a[i] {
				if b == 'O' {
					c[0][(i+j)%3]++
				} else if b == 'X' {
					c[1][(i+j)%3]++
				}
			}
		}
		miI := 0
		for i := 1; i < 3; i++ {
			if c[0][i]+c[1][(i+1)%3] < c[0][miI]+c[1][(miI+1)%3] {
				miI = i
			}
		}
		for i, r := range a {
			for j, b := range r {
				if b == 'O' && (i+j)%3 == miI {
					r[j] = 'X'
				} else if b == 'X' && (i+j)%3 == (miI+1)%3 {
					r[j] = 'O'
				}
			}
			Fprintln(out, string(r))
		}
	}
}

//func main() { CF1450C2(os.Stdin, os.Stdout) }
