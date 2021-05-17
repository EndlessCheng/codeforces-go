package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF873D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	if k&1 == 0 || k > 2*n-1 {
		Fprint(out, -1)
		return
	}
	a := make([]int, n)
	for i := range a {
		a[i] = i + 1
	}
	var f func([]int)
	f = func(a []int) {
		n := len(a)
		if k == 0 || n == 1 {
			return
		}
		k -= 2
		m := n / 2
		b := append([]int(nil), a...)
		copy(a, b[n-m:])
		copy(a[m:], b)
		f(a[:m])
		f(a[m:])
	}
	k--
	f(a)
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

//func main() { CF873D(os.Stdin, os.Stdout) }
