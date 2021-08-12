package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1178D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	isPrime := func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	var n int
	Fscan(in, &n)
	m := n
	for ; !isPrime(m); m++ {
	}
	Fprintln(out, m)
	for i := 1; i < n; i++ {
		Fprintln(out, i, i+1)
	}
	Fprintln(out, n, 1)
	m -= n
	if m == 1 {
		Fprintln(out, 1, 3)
	} else {
		for i := 1; i <= m; i++ {
			Fprintln(out, i, i+m)
		}
	}
}

//func main() { CF1178D(os.Stdin, os.Stdout) }
