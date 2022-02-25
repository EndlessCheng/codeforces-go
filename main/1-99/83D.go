package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF83D(in io.Reader, out io.Writer) {
	var l, r, k int
	Fscan(in, &l, &r, &k)
	for i := 2; i*i <= k; i++ {
		if k%i == 0 {
			Fprint(out, 0)
			return
		}
	}
	if r/k < k {
		if l <= k && k <= r {
			Fprint(out, 1)
		} else {
			Fprint(out, 0)
		}
		return
	}
	l--

	primes := []int{}
	np := make([]bool, k)
	for i := 2; i < k; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := 2 * i; j < k; j += i {
				np[j] = true
			}
		}
	}
	var f func(int, int) int
	f = func(i, v int) int {
		ans := r/v - l/v
		for j := i; j < len(primes) && primes[j] <= r/v; j++ {
			ans -= f(j+1, v*primes[j])
		}
		return ans
	}
	Fprint(out, f(0, k))
}

//func main() { CF83D(os.Stdin, os.Stdout) }
