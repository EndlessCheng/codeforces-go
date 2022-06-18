package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1280B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]string, n)
		tot, fullR := 0, false
		col := make([]int, m)
		for i := range a {
			Fscan(in, &a[i])
			r := 0
			for j, c := range a[i] {
				if c == 'A' {
					r++
					col[j]++
				}
			}
			tot += r
			if r == m {
				fullR = true
			}
		}
		fullC := false
		for _, c := range col {
			if c == n {
				fullC = true
			}
		}
		r, rr := strings.Count(a[0], "A"), strings.Count(a[n-1], "A")
		switch {
		case tot == 0:
			Fprintln(out, "MORTAL")
		case tot == n*m:
			Fprintln(out, 0)
		case r == m || rr == m || col[0] == n || col[m-1] == n:
			Fprintln(out, 1)
		case fullR || fullC || a[0][0] == 'A' || a[0][m-1] == 'A' || a[n-1][0] == 'A' || a[n-1][m-1] == 'A':
			Fprintln(out, 2)
		case r > 0 || rr > 0 || col[0] > 0 || col[m-1] > 0:
			Fprintln(out, 3)
		default:
			Fprintln(out, 4)
		}
	}
}

//func main() { CF1280B(os.Stdin, os.Stdout) }
