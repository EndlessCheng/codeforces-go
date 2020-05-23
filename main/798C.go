package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF798C(_r io.Reader, _w io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n, c, g int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		g = gcd(g, a[i])
		a[i] &= 1
	}
	if g == 1 {
		for i, v := range a {
			if v == 1 {
				if i+1 == n {
					c += 2
				} else if a[i+1] == 0 {
					c += 2
				} else {
					a[i+1] = 0
					c++
				}
			}
		}
	}
	Fprintln(_w, "YES")
	Fprint(_w, c)
}

//func main() { CF798C(os.Stdin, os.Stdout) }
