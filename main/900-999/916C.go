package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF916C(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	var n, m int
	Fscan(_r, &n, &m)
	sp := n - 1
	for ; !isPrime(sp); sp++ {
	}
	Fprintln(out, sp, sp)
	Fprintln(out, 1, 2, 1+sp-(n-1))
	for i := 2; i < n; i++ {
		Fprintln(out, i, i+1, 1)
	}
	m -= n - 1
	for i := 1; i < n; i++ {
		for j := i + 2; j <= n; j++ {
			if m == 0 {
				return
			}
			m--
			Fprintln(out, i, j, int(1e9))
		}
	}
}

//func main() { CF916C(os.Stdin, os.Stdout) }
