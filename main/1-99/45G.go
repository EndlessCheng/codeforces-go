package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF45G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	isPrime := func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true // n > 1
	}

	var n int
	Fscan(in, &n)
	c := make([]int, n)
	n = n * (n + 1) / 2
	if !isPrime(n) {
		if n&1 > 0 && !isPrime(n-2) {
			c[2] = 2
			n -= 3
		}
		i := 2
		for ; !isPrime(i) || !isPrime(n-i); i++ {
		}
		c[i-1] = 1
	}
	for _, v := range c {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF45G(os.Stdin, os.Stdout) }
