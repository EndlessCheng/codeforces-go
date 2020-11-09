package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF827B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, k int
	Fscan(bufio.NewReader(in), &n, &k)
	l, r := (n-1)/k, (n-1)%k
	Fprintln(out, 2*l+min(r, 2))
	c := k + 2
	for i := 1; i <= k; i++ {
		ll := l
		if i > r {
			ll--
		}
		Fprintln(out, k+1, i)
		if ll > 0 {
			Fprintln(out, i, c)
			c++
		}
		for j := 1; j < ll; j++ {
			Fprintln(out, c-1, c)
			c++
		}
	}
}

//func main() { CF827B(os.Stdin, os.Stdout) }
