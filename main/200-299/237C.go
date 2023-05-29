package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF237C(in io.Reader, out io.Writer) {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var a, b, k, ans int
	Fscan(in, &a, &b, &k)
	primes := []int{a - 1}
	np := make([]bool, b+1)
	for i := 2; i <= b; i++ {
		if !np[i] {
			if i >= a {
				primes = append(primes, i)
			}
			for j := i * i; j <= b; j += i {
				np[j] = true
			}
		}
	}
	primes = append(primes, b+1)

	m := len(primes)
	if m-2 < k {
		Fprint(out, -1)
		return
	}
	for i := k; i < m; i++ {
		ans = max(ans, primes[i]-primes[i-k])
	}
	Fprint(out, ans)
}

//func main() { CF237C(os.Stdin, os.Stdout) }
