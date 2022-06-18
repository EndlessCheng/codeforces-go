package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF847E(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	var s string
	Fscan(bufio.NewReader(in), &n, &s)
	var a, b []int
	for i, c := range s {
		if c == 'P' {
			a = append(a, i)
		} else if c == '*' {
			b = append(b, i)
		}
	}
	Fprint(out, sort.Search(n*2, func(t int) bool {
		i := 0
		for _, x := range a {
			y := b[i]
			if x-y > t {
				return false
			}
			if y-x > t {
				continue
			}
			r := x + t
			if y < x {
				r = max(t+2*y-x, (t+y+x)/2)
			}
			for b[i] <= r {
				if i++; i == len(b) {
					return true
				}
			}
		}
		return false
	}))
}

//func main() { CF847E(os.Stdin, os.Stdout) }
