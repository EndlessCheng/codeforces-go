package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1009E(_r io.Reader, out io.Writer) {
	const mod = 998244353
	pow2 := [1e6]int64{1}
	for i := 1; i < 1e6; i++ {
		pow2[i] = pow2[i-1] << 1 % mod
	}
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	c := pow2[n-1] * int64(a[0])
	s := int64(0)
	for i, v := range a {
		s += c
		if i+1 < n {
			c = (c + pow2[n-2-i]*int64(a[i+1]-v)) % mod
		}
	}
	Fprint(out, s%mod)
}

//func main() { CF1009E(os.Stdin, os.Stdout) }
