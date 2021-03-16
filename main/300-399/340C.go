package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF340C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, s, pre int64
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for i, v := range a {
		v := int64(v)
		s += (int64(i)*v - pre) * 2
		pre += v
	}
	s += pre
	g := gcd(s, n)
	Fprint(out, s/g, n/g)
}

//func main() { CF340C(os.Stdin, os.Stdout) }
